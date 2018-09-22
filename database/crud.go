package db

import (
	"fmt"
	// "encoding/json"

	"github.com/elgs/jsonql"
)

// Put is a Collection method that allows you to insert an item
// into a collection
func (c Collection) Put (data interface{}) (interface{}, error) {
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

func Insert(data interface{}, coll string) error {
	c := LoadCollection(coll)
	var doc, err = NewDocument(data)
	if err != nil {
		return Error.New("Error creating new document", err)
	}
	
	c.Data[doc.ID] = doc.Data

	err = SaveCollection(coll, c)
	if err != nil {
		return Error.New("error saving collection", err)
	}
	return nil
}

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

func (c Collection) FindById(id string) (interface{}, bool) {
	if val, ok := c.Data[id]; ok {
		return val, true
	}
	return nil, false
}

func (c Collection) Delete(key string) (interface{}, error) {
	delete(c.Data, key)
	if val, ok := c.Data[key]; ok {
		return val, nil
	}
	return nil, Error.New("error deleting key %+v\n", key)
}
