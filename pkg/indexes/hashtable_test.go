package indexes

import (
	"fmt"
	"os"
	"testing"
)

func TestHashTables(t *testing.T) {
	t.Run("Test HashTable Get", func(t *testing.T) {
		tmpDir := os.TempDir()
		defer os.Remove(tmpDir)

		ht, err := NewHashTableIndex(tmpDir)
		if err != nil {
			t.Fail()
		}

		inserts := []struct {
			key     string
			value   interface{}
			wantErr error
		}{
			{
				key:     "test happy path",
				value:   "happy path tested",
				wantErr: nil,
			},
			{
				key:     "test bytes",
				value:   []byte("fuck yeah bytes"),
				wantErr: nil,
			},
		}

		for _, i := range inserts {
			fmt.Printf("i.key %+v", i.key)
			fmt.Printf("i.value %+v", i.value)
			ht.Put(i.key, i.value)
		}

		getCases := []struct {
			key      string
			expected interface{}
			wantErr  error
		}{
			{
				expected: nil,
			},
		}

		for _, i := range getCases {
			val := ht.Get(i.key)
			if val != i.expected {
				t.Fail()
			}
		}
	})
}
