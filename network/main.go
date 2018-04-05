package main

import (
	"time"
	"os"
	"./net"
	"./consensus"
	"fmt"
	"os/signal"
)

func main() {
	// Initialize
	network := net.Bootstrap()
	node := network[0]

	// Main loop status channel
	program := make(chan int)

	// Kick off gossip loop
	interval := time.Second / 5
	go consensus.Run(program, node, interval)

	// Block until `program` is closed or SIGINT
	go handleSignals(program)
	os.Exit(block(program, 0))
}

func block(program chan int, status int) int {
	_status, ok := <-program
	if !ok {
		fmt.Println("last status:", _status)
		return status
	}

	return block(program, _status)
}

func handleSignals(program chan int) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for range interrupt {
		fmt.Printf("Stopping... ")
		close(program)
	}
}
