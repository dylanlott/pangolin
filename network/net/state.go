package net

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type TxId int

type Tx struct {
	Id TxId
}

type State struct {
	Db     *bolt.DB
	Bucket []byte
}

type Diff struct {
	state1 *State
	state2 *State
	equal  bool
	data   []Tx
}

// TODO: return *[]TxId, *[]Tx
func (state *State) read(state2 *State) (*[]string, *[]string) {
	keys, values := new([]string), new([]string)
	state.Db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(state.Bucket).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			*keys = append(*keys, string(k))
			*values = append(*values, string(v))
			// fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	return keys, values
}

func (state *State) diff(state2 *State) (Diff, Diff) {
	keys, values := state.read(state2)

	fmt.Println(keys, values)
	diff1 := Diff{equal: true}
	diff2 := Diff{equal: true}
	// diff1 := new(Diff)
	// diff2 := new(Diff)

	// ...

	return diff1, diff2
}

func (state *State) write(diff *Diff) {
	if diff.equal {
		return
	}

	fmt.Printf("updating State: %v\nwith diff: %v", state, diff)
}
