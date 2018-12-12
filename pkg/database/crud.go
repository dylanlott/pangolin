package db

import (
	"encoding/json"
	"fmt"

	"github.com/elgs/jsonql"
	"github.com/thedevsaddam/gojsonq"
)

// Put is a Collection method that allows you to insert an item
// into a collection
func (c Collection) Put(data interface{}) (interface{}, error) {
	doc, err := NewDocument(data)
	if err != nil {
		return nil, err
	}
	c.Data[doc.ID] = doc.Data
	err = SaveCollection(c.Name, c)
	if err != nil {
		return nil, Error.New("Error saving collection after Put: %+v\n", err)
	}
	return c.Data[doc.ID], nil
}

// Insert takes a data interface and a collection and inserts it
// into that collection, returning the inserted data and an error
func Insert(data interface{}, coll string) (interface{}, error) {
	c := LoadCollection(coll)
	var doc, err = NewDocument(data)
	if err != nil {
		return nil, Error.New("error creating new document", err)
	}

	c.Data[doc.ID] = doc.Data

	err = SaveCollection(coll, c)
	if err != nil {
		return nil, Error.New("error saving collection", err)
	}
	return doc, nil
}

// NewQuery creates a new Query object for chain use elsewhere
func (c Collection) NewQuery() *gojsonq.JSONQ {
	data, err := json.Marshal(c.Data)
	if err != nil {
		fmt.Printf("error marshaling JSON %+v\n", err)
		return nil
	}
	return gojsonq.New().JSONString(string(data))
}

// Where returns a Where query set that matches key, op, val parameters
func (c Collection) Where(key, op, val string) interface{} {
	q := c.NewQuery().Where(key, op, val).Get()
	return q
}

// Find is for querying the JSON of collections
func (c Collection) Find(q string) error {
	parser := jsonql.NewQuery(c.Data)
	data, err := parser.Query(q)
	if err != nil {
		fmt.Printf("error parsing query: %+v\n", err)
		return err
	}
	fmt.Printf("FOUND DATA: %+v\n", data)
	return nil
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// FindOne finds exactly one collection.
func (c Collection) FindOne(q string) (interface{}, error) {
	parser := jsonql.NewQuery(c.Data)
	data, err := parser.Query(q)
	if err != nil {
		fmt.Printf("Error FindOne: %+v\n", err)
		return nil, err
	}
	fmt.Printf("data retrieved %+v\n", data)
	return data, err
}

// Delete takes a key and deletes the document located at that key
func Delete(key string, coll string) (interface{}, bool) {
	c, err := GetCollection(coll)
	checkError(err)

	snapshot := c.Data[key]
	delete(c.Data, key)

	_, ok := c.Data[key]
	if ok {
		err = SaveCollection(coll, c)
		checkError(err)
		return snapshot, true
	}

	return nil, false
}

// Delete is a collection method that deletes the document at argument
// `key` and returns an interface and error
func (c Collection) Delete(key string) (interface{}, error) {
	delete(c.Data, key)
	if val, ok := c.Data[key]; ok {
		return val, nil
	}
	err := SaveCollection(c.Name, c)
	if err != nil {
		return nil, Error.New("Error saving collection after deletion: %+v\n", err)
	}
	return nil, Error.New("error deleting key %+v\n", key)
}
