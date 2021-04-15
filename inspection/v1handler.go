package inspection

import (
	"context"
	"errors"
	"github.com/ipfs/go-ipfs/miner/proto"
	"github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	"ipfc/dbstore/ds"
	"ipfc/dbstore/model"
	"ipfc/subpub"
	"time"
)

type V1Handler struct {
	api        iface.CoreAPI
	store      *ds.DbStore
	handleFunc map[string]subpub.HandleFunc
}

func NewV1Handler(api iface.CoreAPI, store *ds.DbStore) *V1Handler {
	h := &V1Handler{
		api:        api,
		store:      store,
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
		return errors.New("proto error")
	}
	//todo: check WalletAddress
	if hbMsg.WalletAddress == "" {
		return errors.New("WalletAddress is empty")
	}
	if hbMsg.Role != model.RoleMainNode && hbMsg.Role != model.RoleEdgeNode {
		return errors.New("miner role error")
	}
	miner := &model.Miner{
		Id:           receivedFrom.String(),
		Address:      hbMsg.WalletAddress,
		Role:         hbMsg.Role,
		CreatedAt:    time.Now().Unix(),
		LastActiveAt: time.Now().Unix(),
	}
	if h.store.MinerExist(miner.Id) {
		return h.store.UpdateMiner(miner)
	}
	return h.store.CreateMiner(miner)
}

func (h *V1Handler) HandleWindowPostResp(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error {
	//todo:
	log.Infof("HandleWindowPostResp: %+v", msg)
	return nil
}
