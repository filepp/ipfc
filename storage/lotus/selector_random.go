package lotus

import (
	"context"
	"ipfc/storage/types"
	"math/rand"
)

type RandomMinerSelector struct {
}

func NewRandomMinerSelector() *RandomMinerSelector {
	return &RandomMinerSelector{}
}

func (s *RandomMinerSelector) GetMiners(ctx context.Context, fileInfo *types.FileInfo, count int, filter types.MinerSelectorFilter) ([]string, error) {
	candidates := filter(ctx, fileInfo)
	if len(candidates) <= count {
		return candidates, nil
	}

	selectedMinerMark := make(map[int]struct{})
	selectedMiners := make([]string, 0)

	for n := count; n > 0; {
		index := rand.Int() % len(candidates)
		if _, ok := selectedMinerMark[index]; ok {
			continue
		}
		selectedMinerMark[index] = struct{}{}
		selectedMiners = append(selectedMiners, candidates[index])
		n--
	}
	return selectedMiners, nil
}

var _ types.MinerSelector = &RandomMinerSelector{}
