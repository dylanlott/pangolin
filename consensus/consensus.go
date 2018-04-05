package main

import (
	"time"
	"os"
	"fmt"
	"os/signal"
	"../net"
)

func main() {
	// _, err := bolt.Open("bolt.db", 0644, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	network := net.Bootstrap()
	node := network[0]

	program := make(chan int)
	go consensus(program, node)
	go handleSignals(program)
	os.Exit(run(program, 0))
}

func consensus(program chan int, node *net.NetNode) {
	ticker := time.NewTicker(time.Second / 5)

	// i := 0
	for t := range ticker.C {
		node.RandomGossip(t)
	}
}

func run(program chan int, status int) int {
	_status, ok := <-program
	if !ok {
		fmt.Println("last status:", _status)
		return status
	}

	return run(program, _status)
}

func handleSignals(program chan int) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for range interrupt {
		fmt.Printf("Stopping... ")
		close(program)
	}
}
