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
}
