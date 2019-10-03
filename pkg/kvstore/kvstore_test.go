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

	kv, err := BadgerKVStore(testDir)
	assert.NoError(t, err)
	assert.NotNil(t, kv)
	assert.NotNil(t, kv.db)

	t.Run("test kvstore", func(t *testing.T) {
		// happy path
		key := []byte("testkey")
		value := []byte("test value")

		err := kv.Put(key, value)
		t.Log(err)
		assert.NoError(t, err)

		val, err := kv.Get(key)
		assert.NoError(t, err)
		assert.NotNil(t, val)

		// getting non existent key should not error, only return nil
		val, err = kv.Get([]byte("nonexistent"))
		assert.NoError(t, err)
		assert.Nil(t, val)

		// test that nil keys throw errors
		err = kv.Put([]byte(""), []byte("test empty value"))
		assert.Error(t, err)

		// test that nil values throw errors
		err = kv.Put([]byte("testkey"), nil)
		assert.Error(t, err)
	})
}
