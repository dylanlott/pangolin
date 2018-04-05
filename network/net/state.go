package net

type State struct {
	keys []string
	values []string
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