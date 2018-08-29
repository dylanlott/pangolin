package main

import (
	"fmt"

	"./database"
)

func main() {
	p, err := db.NewDatabase()
	fmt.Printf("p: %+v\n", p)
	fmt.Printf("err: %+v\n", err)
	return
}
