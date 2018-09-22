package main

import (
	"fmt"

	"github.com/dylanlott/pangolin/database"
)

func main() {
	err := db.NewDatabase()
	if err != nil {
		fmt.Printf("Error setting up database: %+v\n", err)
	}

	Get()
}

func InsertData() {
	coll, err := db.GetCollection("name")
	fmt.Printf("Got Collection: %+v\n", coll)

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

func Get() {
	coll, err := db.GetCollection("name")	
	if err != nil {
		fmt.Printf("ERROR getting collection %+v\n", err)
	}
	coll.Find("integer > 1")
}
