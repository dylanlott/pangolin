package net

import "github.com/boltdb/bolt"

var db, Err = bolt.Open("bolt.db", 0644, nil)

type State struct {
}

type Diff struct {
	state1 *State
	state2 *State
}

func (s *State) diff(otherState State) (Diff, Diff) {
	return Diff{}, Diff{}
}

func (s *State) apply(diff Diff) {

}