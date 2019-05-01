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
type Value map[string]interface{}

// String returns the value of a Key as a string
func (k Key) String() string {
	return string(k)
}

// NewPair returns a Pair struct that adheres to JSON. If it does not adhere
// to JSON, it will return an error
func NewPair(key string, value interface{}) (pangolin.Pair, error) {
	var key []byte
	var val []byte

	// validate data for JSON compatibility

	// TODO: Marshall values to bytes
	// TODO: Probably should use protobuf here - https://godoc.org/github.com/golang/protobuf/proto#Marshal
	return pangolin.Pair{
		Key:   pangolin.Key,
		Value: pangolin.Value,
	}
}
