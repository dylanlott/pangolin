package db

import (
	"io"
	"os"
	"sync"
	"encoding/json"
	"bytes"
	"fmt"
	"crypto/rand"
)

var lock sync.Mutex

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
  lock.Lock()
  defer lock.Unlock()
  f, err := os.Create(path)
  if err != nil {
    return err
  }
  defer f.Close()
  r, err := Marshal(v)
  if err != nil {
    return err
  }
  _, err = io.Copy(f, r)
  return err
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
  lock.Lock()
  defer lock.Unlock()
  f, err := os.Open(path)
  if err != nil {
    return err
  }
  defer f.Close()
  return Unmarshal(f, v)
}

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
  b, err := json.MarshalIndent(v, "", "\t")
  if err != nil {
    return nil, err
  }
  return bytes.NewReader(b), nil
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
  return json.NewDecoder(r).Decode(v)
}

// Creates a new UUID as a unique identifier for each item 
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
