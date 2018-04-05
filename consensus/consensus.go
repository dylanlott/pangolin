// package consensus
package main

import (
	"time"
	"os"
	"fmt"
)

func main() {
	// _, err := bolt.Open("bolt.db", 0644, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	program := make(chan int)
	go consensus(program)
	os.Exit(run(program, 0))
}

func consensus(program chan int) {
	ticker := time.NewTicker(time.Second / 5)

	i := 0
	for t := range ticker.C {
		fmt.Println("ticker at:", t)
		i++
		switch {
		case i > 10:
			close(program)
		case i > 5:
			program <- 100 + i
		}
	}
}

func run(program chan int, status int) int {
	_status, ok := <-program
	fmt.Println("Status updated:", _status)
	if !ok {
		return status
	}

	return run(program, _status)
}
