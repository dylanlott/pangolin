package net

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type State struct {
	db bolt.DB
	bucket bolt.Bucket
}

type Diff struct {
	state1 *State
	state2 *State
	equal bool
	data State
}

func (s *State) diffKeys(state *State) ([]string, []int) {
	keys := make([]string, 0)
	indexes := make([]int, 0)

	// consensus.Db.
	//
	// for i, key := range s.keys {
	// 	state.keys[i]
	// }
	return keys, indexes
}

func (s *State) diff(state *State) (Diff, Diff) {
	// keys, indexes := s.diffKeys(state)
	diff1 := Diff{equal: true}
	diff2 := Diff{equal: true}
	// diff1 := new(Diff)
	// diff2 := new(Diff)

	// ...

	return diff1, diff2
}

func (s *State) apply(diff *Diff) {
	if diff.equal {
		return
	}

	fmt.Printf("updating s: %v\nwith diff: %v", s, diff)
}
