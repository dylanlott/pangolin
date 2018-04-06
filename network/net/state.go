package net

import (
	"fmt"
	"github.com/boltdb/bolt"
)

// type TxId int
//
// type Tx struct {
// 	Id TxId
// }
type Any interface{}
type Keys []Any
type Values []Any

type State struct {
	Db     *bolt.DB
	Bucket []byte
}

type Diff struct {
	state1 State
	state2 State
	data   map[Any]Any
	// data   []Tx
}

// TODO: return *[]TxId, *[]Tx
func (state *State) read(state2 *State) (Keys, Values) {
	keys, values := new(Keys), new(Values)

	state.Db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(state.Bucket).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			*keys = append(*keys, string(k))
			*values = append(*values, string(v))
			// fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	return *keys, *values
}

func (state *State) diff(state2 *State) (diff, diff2 Diff) {
	diff = Diff{*state, *state2, &map[Any]Any{}}
	diff2 = Diff{*state2, *state,}
	keys, values := state.read(state)
	keys2, values2 := state.read(state2)

	for uniqueKey := range sliceDifference(keys, keys2) {

	}
	// NB: don't do this
	// valueDiff := sliceDifference(values, values2)

	fmt.Println(keys, values)

	// ...

	return diff, diff2
}

func (state *State) write(diff *Diff) {
	// if diff.equal {
	// 	return
	// }

	fmt.Printf("updating State: %v\nwith diff: %v", state, diff)
}

type elem struct {
	index  int
	index2 int
	value  Any
	unique bool
}

func sliceDifference(slice, slice2 []Any) (diff []elem) {
	m := map[Any]elem{}

	for i, v := range slice {
		m[v] = elem{i, nil, v, true}
	}
	for i, v := range slice2 {
		e, ok := m[v]
		if ok {
			e.unique = false
			e.index2 = i
			break
		}

		m[v] = elem{nil, i, v, true}
	}

	for _, e := range m {
		if e.unique == true {
			diff = append(diff, e)
		}
	}

	return diff
}
