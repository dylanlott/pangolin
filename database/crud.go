package db

import (
	"fmt"
	"log"
)

func Insert(data interface{}, coll string) error {
	// get snapshot of collection 
	c := LoadCollection(coll)
	log.Printf("got collection %+v\n", c)

	// create document
	var doc, err = NewDocument(data)

	if err != nil {
		return Error.New("Error creating new document", err)
	}
	log.Printf("saving document %+v\n", doc)
	log.Printf("previous data: %+v\n", c.Data)
	
	// save document
	c.Data = append(c.Data, doc)

	log.Printf("added doc to collection %+v\n", c.Data)
	// if no errors, complete transaction

	err = SaveCollection(coll, c)
	if err != nil {
		fmt.Printf("error saving collection %+v\n", err)
		return Error.New("error saving collection", err)
	}
	return nil
}

func Find() {

	// Load collection

	// Parse as json

	// Parse query on json 

	// get results 

	// format as response and return
}

func FindOne() {

	// Load collection

	// Find single by query or id

	// If multiple, return first object in list

	// Create response for document and return it
}

func Delete() {
	// Load Collection

	// keep collection as snapshot 

	// Find Documents by query with parser 

	// Remove documents from collection 

	// Save collection 

	// If error saving, revert back to snapshot of collection

	// Return documents if no error

	// Return error if transaction did not occur
}
