package kvstore

// Key is a string for keys
type Key []byte

// Value is a marshalled byte slice of the Value
type Value []byte

// Keyed outlines what a Key has to be able to return
type Keyed interface {
	Bytes() []byte
	String() string
}

// Paired is the interface that Pair must fulfill
type Paired interface {
	Key() (Keyed, error)
	Value() (Valued, error)
}

// Valued standardizes all storage values to an interface
type Valued interface {
	Bytes() []byte
	Value() interface{}
}

// KeyValueStore is the main interface that any driver
// must fulfill to be a driver for the DB.
type KeyValueStore interface {
	Get(key Keyed) error
	Put(key Keyed, value Valued) error
	Delete(key Keyed) error
}

// Database is the entire wrapper of KeyValueStore, Indexes,
// and query engines.
type Database interface {
	KVStore() *KeyValueStore
	Indexes() []string
	Query(query string) interface{}
}

// Pair fulfills the Paired interface.
type Pair struct {
	Key   Keyed
	Value Valued
}

// NewPair returns a Pair that fulfills Paired, Keyed, and Valued
func NewPair(key []byte, value []byte) Pair {
	return Pair{}
}

// Put acts on a Pair struct and inserts them into the KVStore
func (p Pair) Put() error {
	return nil
}

// Get acts on a Pair and returns the value of the Key in that Pair
func (p Pair) Get() (Valued, error) {
	return nil, nil
}
