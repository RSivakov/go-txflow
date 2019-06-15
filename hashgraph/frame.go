package hashgraph

import (
	"sort"

	"github.com/tendermint/tendermint/types"
)

type Frame struct {
	Round         int //RoundReceived
	Validators    []*types.Validator
	Roots         map[string]*Root
	Events        []*FrameEvent              //Events with RoundReceived = Round
	ValidatorSets map[int][]*types.Validator //[round] => Peers
}

func (f *Frame) SortedFrameEvents() []*FrameEvent {
	sorted := SortedFrameEvents{}
	for _, r := range f.Roots {
		sorted = append(sorted, r.Events...)
	}
	sorted = append(sorted, f.Events...)
	sort.Sort(sorted)
	return sorted
}
