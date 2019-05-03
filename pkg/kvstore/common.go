package kvstore

import (
	"/pkg/pangolin"

	"github.com/zeebo/errs"
)

// KVStore is the interface for all KV Stores we accept. It's purposefully broad to push type
// handling and logic into the driver for each KV Store that fulfills this interface.
// type KVStore interface {
// 	Get(key pangolin.Key) (pangolin.Pair, bool, error)
// 	Put(key pangolin.Key, val pangolin.Value) error
// 	Delete(key pangolin.Key) error
// }

var (
	// Error is the main error class for this package
	Error = errs.Class("kvstore error")
)

// Get takes a KVStore and a key and returns the value or an error from that key
func Get(k *KVStore, key string) (*pangolin.Pair, bool, error) {
	return k.Get(key)
}

// Put takes a KV Store and puts a value to it
func Put(k *KVStore, key pangolin.Key, value pangolin.Value) error {
	return k.Put(key, value)
}

// Delete takes a KV store and deletes the value with key `k`. It returns an error if it was unsuccessful.
func Delete(k *KVStore, key pangolin.Key) error {
	return k.Delete(key)
}
