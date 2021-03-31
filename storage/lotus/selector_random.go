package lotus

import (
	"context"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"ipfc/storage/types"
	"math/rand"
)

type RandomMinerSelector struct {
	minerAsks map[string]*storagemarket.StorageAsk
	count     int
}

func NewRandomMinerSelector(minerAsks map[string]*storagemarket.StorageAsk, count int) *RandomMinerSelector {
	return &RandomMinerSelector{
		minerAsks: minerAsks,
		count:     count,
	}
}

func (s *RandomMinerSelector) GetMiners(ctx context.Context, fileInfo *types.FileInfo) ([]string, error) {
	candidates := s.getCandidateMiner(ctx, fileInfo)
	if len(candidates) <= s.count {
		return candidates, nil
	}

	selectedMinerMark := make(map[int]struct{})
	selectedMiners := make([]string, 0)

	for count := s.count; count > 0; {
		index := rand.Int() % len(candidates)
		if _, ok := selectedMinerMark[index]; ok {
			continue
		}
		selectedMinerMark[index] = struct{}{}
		selectedMiners = append(selectedMiners, candidates[index])
		count--
	}
	return selectedMiners, nil
}

func (s *RandomMinerSelector) getCandidateMiner(ctx context.Context, fileInfo *types.FileInfo) []string {
	var miners []string
	for miner, minerAsk := range s.minerAsks {
		if fileInfo.Size < uint64(minerAsk.MinPieceSize) || fileInfo.Size > uint64(minerAsk.MaxPieceSize) {
			continue
		}
		miners = append(miners, miner)
	}
	return miners
}

var _ types.MinerSelector = &RandomMinerSelector{}
