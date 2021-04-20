package inspection

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ipfs/go-cid"
	ds2 "github.com/ipfs/go-datastore"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/miner/proto"
	"github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/libp2p/go-libp2p-core/peer"
	"io"
	"ipfc/dbstore/ds"
	"ipfc/dbstore/model"
	"ipfc/subpub"
	"time"
)

type V1Handler struct {
	api        iface.CoreAPI
	store      *ds.DbStore
	localStore ds2.Datastore
	handleFunc map[string]subpub.HandleFunc
}

func NewV1Handler(api iface.CoreAPI, store *ds.DbStore, localStore ds2.Datastore) *V1Handler {
	h := &V1Handler{
		api:        api,
		store:      store,
		localStore: localStore,
		handleFunc: make(map[string]subpub.HandleFunc),
	}
	h.handleFunc[proto.MsgMinerHeartBeat] = h.HandleMinerHeartBeat
	h.handleFunc[proto.MsgWindowPostResponse] = h.HandleWindowPostResp
	return h
}

func (h *V1Handler) Handle(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error {
	if f, ok := h.handleFunc[msg.Type]; ok {
		return f(ctx, receivedFrom, msg)
	}
	log.Warnf("message type not register: %v", msg.Type)
	return nil
}

func (h *V1Handler) HandleMinerHeartBeat(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error {
	log.Infof("HandleMinerHeartBeat: %v,  %+v", receivedFrom.String(), msg)
	hbMsg, ok := msg.Data.(proto.MinerHartBeat)
	if !ok {
		log.Errorf("proto error")
		return errors.New("proto error")
	}
	if hbMsg.Role != model.RoleMainNode && hbMsg.Role != model.RoleEdgeNode {
		return errors.New("miner role error")
	}
	miner := &model.Miner{
		Id:           receivedFrom.String(),
		Role:         hbMsg.Role,
		CreatedAt:    time.Now().Unix(),
		LastActiveAt: time.Now().Unix(),
	}
	if h.store.MinerExist(miner.Id) {
		return h.store.UpdateMiner(miner)
	}
	return nil
}

func (h *V1Handler) HandleWindowPostResp(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error {
	log.Infof("HandleWindowPostResp: %+v", msg)

	wndPostResp, ok := msg.Data.(proto.WindowPostResp)
	if !ok {
		log.Errorf("proto error")
		return errors.New("proto error")
	}
	wndPostReq, err := getWindowPostMessage(h.localStore, receivedFrom.String(), msg.Nonce)
	if err != nil {
		log.Errorf("failed to get wndPostMsg: %v, %v %v", err, receivedFrom.String(), msg.Nonce)
		return err
	}
	if len(wndPostReq.Items) == 0 {
		return nil
	}
	passRate, err := h.verifyWindowPost(ctx, &wndPostReq, &wndPostResp)
	if err != nil {
		return err
	}
	wnd := &model.WindowPost{
		MinerId:   receivedFrom.String(),
		PassRate:  passRate,
		CreatedAt: time.Now().Unix(),
	}
	reqdata, _ := json.Marshal(wndPostReq)
	resultDdata, _ := json.Marshal(wndPostResp)
	wnd.Request = string(reqdata)
	wnd.Result = string(resultDdata)

	h.store.SaveWindowPost(wnd)

	deleteWindowPostMessage(h.localStore, receivedFrom.String(), msg.Nonce)
	return nil
}

func (h *V1Handler) verifyWindowPost(ctx context.Context, wndPostReq *proto.WindowPostReq, wndPostResp *proto.WindowPostResp) (float64, error) {
	reqMap := make(map[string]proto.WindowPostReqItem)
	for _, v := range wndPostReq.Items {
		reqMap[v.FileCid.String()] = v
	}
	successCount := 0
	for _, v := range wndPostResp.Items {
		if req, ok := reqMap[v.FileCid.String()]; ok {
			if h.verifyPostOfFile(ctx, &req, &v) {
				successCount++
			}
		}
	}
	return float64(successCount) / float64(len(wndPostReq.Items)), nil
}

func (h *V1Handler) verifyPostOfFile(ctx context.Context, wndReq *proto.WindowPostReqItem, wndResp *proto.WindowPostRespItem) bool {
	if len(wndReq.Positions) != len(wndResp.Positions) || len(wndResp.Positions) != len(wndResp.Data) {
		return false
	}
	data, err := h.getFileDataAtFixedPosition(ctx, wndReq.FileCid, wndReq.Positions)
	if err != nil {
		log.Errorf("failed to get file data: %v", err)
		return false
	}
	for i, v := range data {
		if v != wndResp.Data[i] {
			return false
		}
	}
	log.Infof("verify post of file success: %v ", wndReq.FileCid.String())
	return true
}

func (h *V1Handler) getFileDataAtFixedPosition(ctx context.Context, fileCid cid.Cid, positions []int64) ([]byte, error) {
	cidPath := path.New("/ipfs/" + fileCid.String())

	fileNode, err := h.api.Unixfs().Get(ctx, cidPath)
	if err != nil {
		log.Errorf("failed to get:%v", err.Error())
		return nil, err
	}
	defer fileNode.Close()
	file, ok := fileNode.(files.File)
	if !ok {
		log.Warnf("file type error")
		return nil, errors.New("file type error")
	}

	size, _ := file.Size()
	data := make([]byte, len(positions))
	for i, pos := range positions {
		if pos < 0 || pos >= size {
			log.Warnf("out of range: %v %v %v", fileCid.String(), size, pos)
			return data, errors.New("out off range")
		}
		file.Seek(pos, io.SeekStart)
		var buf [1]byte
		file.Read(buf[:])
		data[i] = buf[0]
	}
	return data, nil
}
