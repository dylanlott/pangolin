package main

import (
	"fmt"

	"github.com/derekparker/trie"
	"github.com/dylanlott/pangolin/database"
)

func main() {
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

	return
}
