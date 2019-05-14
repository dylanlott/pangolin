package pangolin

// Reference: https://blog.golang.org/json-and-go

// Pair is a type for our KV to digest. All KV stores that implement
// and work with the Pair type for Put operations.
type Pair struct {
	Key   Key
	Value Value
}

// Key is the main type for Keys in our KV Store.
type Key []byte

// Value is the main type for values in our KV Store.
// This has to allow an interface to work with JSON data being Marshaled
// and Unmarshaled to byte slices.
type Value map[interface{}]interface{}

// String returns the value of a Key as a string
func (k Key) String() string {
	return string(k)
}

// Bytes returns the bytes of a given Key
func (k Key) Bytes() []byte {
	panic("not impl")
}
