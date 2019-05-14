package indexes

import (
	"testing"

	"github.com/dylanlott/pangolin/pkg/pangolin"
)

func TestCreate(t *testing.T) {
	t.Skip("skip create")
}

func TestGet(t *testing.T) {
	t.Skip("skip get")
}

func TestValueHashTable(t *testing.T) {
	cases := []struct {
		key   string
		value pangolin.Value
		err   error
	}{
		// TODO: add more error path test cases
		{
			key:   "test",
			value: getValue(),
			err:   nil,
		},
	}

	for _, tc := range cases {
		// todo: add hash table creation and removal
		ht := ValueHashtable{}
		ht.Put(tc.key, getValue(tc.Value))
	}
}

func getValue(val interface{}) {
	panic("not impl")
}
