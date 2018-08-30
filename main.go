package main

import (
	"fmt"

	"github.com/dylanlott/pangolin/database"
)

func main() {
	p, err := db.NewDatabase()
	fmt.Printf("p: %+v\n", p)
	fmt.Printf("err: %+v\n", err)
	return
}
