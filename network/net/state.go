package net

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type State struct {
	Db     *bolt.DB
	Bucket []byte
}

type Diff struct {
	state1 *State
	state2 *State
	equal  bool
	data   State
}

func (state *State) diffKeys(state2 *State) ([]string, []int) {
	keys := make([]string, 0)
	indexes := make([]int, 0)

	// state.Db.View(func(tx *bolt.Tx) error {
	// 	c := tx.Bucket(state.bucket).Cursor()
	//
	// 	for k, v := c.First(); k != nil; k, v = c.Next() {
	// 		fmt.Printf("key=%s, value=%s\n", k, v)
	// 	}
	//
	// 	return nil
	// })

	// for i, key := range State.keys {
	// 	state2.keys[i]
	// }
	return keys, indexes
}

func (state *State) diff(state2 *State) (Diff, Diff) {
	keys, indexes := state.diffKeys(state)
	fmt.Println(keys, indexes)
	diff1 := Diff{equal: true}
	diff2 := Diff{equal: true}
	// diff1 := new(Diff)
	// diff2 := new(Diff)

	// ...

	return diff1, diff2
}

func (state *State) apply(diff *Diff) {
	if diff.equal {
		return
	}

	fmt.Printf("updating State: %v\nwith diff: %v", state, diff)
}
