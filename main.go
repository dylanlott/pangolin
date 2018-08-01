package main

import (
	"fmt"
	"log"

	"github.com/dylanlott/pangolin/database"
)

func main() {
	fmt.Printf("starting database")
	pangolin, err := db.NewDatabase()

	if err != nil {
		log.Fatal("Error starting database", err)
		return
	}

	i, err := pangolin.Insert("test", "1234")
	fmt.Printf("inserted correctly %+v\n", i)
	if err != nil {
		fmt.Printf("error %s", err)
	}

	get, err := pangolin.Get("test")
	fmt.Printf("get %+v\n", get)

	if err != nil {
		fmt.Printf("error %s", err)
	}

	number, err := pangolin.Insert("1234", 1234)
	if err != nil {
		fmt.Printf("error inserting integer key", err)
	}

	fmt.Printf("inserted integer key", number)

	getNumber, err := pangolin.Get("1234")
	fmt.Printf("getNumber", getNumber)

	return
}
