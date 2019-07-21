package kvstore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDir string = "./kvtestdata"

func TestKVStore(t *testing.T) {
	err := os.Mkdir(testDir, 0777)
	assert.NoError(t, err)
	defer os.RemoveAll(testDir)

	kv, err := NewBadgerAdapter(testDir)
	assert.NoError(t, err)
	assert.NotNil(t, kv)
}
