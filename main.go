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
	}

	log.Print("created new database")
	return
}
