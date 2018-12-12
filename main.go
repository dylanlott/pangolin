package main

import (
	"fmt"

	"github.com/dylanlott/pangolin/pkg/database"
)

func main() {
	err := db.SetupDatabase()
	if err != nil {
		fmt.Printf("Error setting up database: %+v\n", err)
	}

	// Run this file for exmaples on how to use Pangolin

	// Load a collection
	c := db.LoadCollection("users")

	// Run a Where Query on that Collection
	results := c.Where("6f71ae0f-9efa-4842-81fc-b1d30f86acf3", "<>", "")
	fmt.Printf("results: %+v\n", results)

	// Run your own custom query set on the collection with NewQuery()
	results2 := c.NewQuery().Where("6f71ae0f-9efa-4842-81fc-b1d30f86acf3", "<>", "").Get()
	fmt.Printf("results 2: %+v\n", results2)

	return
}
