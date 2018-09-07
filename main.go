package main

import (
	"log"
	"fmt"

	"github.com/derekparker/trie"
	"github.com/dylanlott/pangolin/database"
)

func main() {
	InsertData()
	LoadCollection()
}

func InsertData() {
	err := db.NewDatabase()
	t := trie.New()
	fmt.Printf("new trie %+v\n", t)

	coll, err := db.NewCollection("name", *t)
	fmt.Printf("Coll: %+v\n", coll)

	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
	}

	data := make(map[string]interface{})
	data["hello"] = "world"
	data["integer"] = 1234
	data["float"] =  1234.56
		
	err = db.Insert(data, coll.Name)

	if err != nil {
		fmt.Printf("Error inserting data %+v\n", err)
	}
}

func LoadCollection() {
	err := db.NewDatabase()
	if err != nil {
		log.Printf("Error creating database: %+v\n", err)
	}
	t := trie.New()
	fmt.Printf("trie %+v\n", t)
	coll := db.LoadCollection("name")	
	log.Printf("Collection: %+v\n", coll)
}
