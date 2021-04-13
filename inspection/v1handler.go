package inspection

import (
	"context"
	"github.com/ipfs/go-ipfs/miner/proto"
	"github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/prometheus/common/log"
	"ipfc/subpub"
)

type V1Handler struct {
	api        iface.CoreAPI
	handleFunc map[string]subpub.HandleFunc
}

func NewV1Handler(api iface.CoreAPI) *V1Handler {
	h := &V1Handler{
		api:        api,
		handleFunc: make(map[string]subpub.HandleFunc),
	}
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

func (h *V1Handler) HandleWindowPostResp(ctx context.Context, receivedFrom peer.ID, msg *proto.Message) error {
	//todo:
	log.Infof("HandleWindowPostResp: %+v", msg)
	return nil
}
