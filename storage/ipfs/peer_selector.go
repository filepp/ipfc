package ipfs

import (
	"context"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"math/rand"
)

type PeerSelector struct {
}

func NewPeerSelector() *PeerSelector {
	return &PeerSelector{}
}

func (s *PeerSelector) GetPeers(ctx context.Context, peers []iface.ConnectionInfo, count int) ([]iface.ConnectionInfo, error) {
	if len(peers) <= count {
		return peers, nil
	}

	selectedMark := make(map[int]struct{})
	selectedPeers := make([]iface.ConnectionInfo, 0)

	for n := count; n > 0; {
		index := rand.Int() % len(peers)
		if _, ok := selectedMark[index]; ok {
			continue
		}
		selectedMark[index] = struct{}{}
		selectedPeers = append(selectedPeers, peers[index])
		n--
	}
	return selectedPeers, nil
}
