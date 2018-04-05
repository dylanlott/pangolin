package main

import (
	"os"
	"./net"
	"./consensus"
	"fmt"
	"os/signal"
	"time"
	"os/user"
	"path/filepath"
)

func main() {
	// Initialize
	u, _ := user.Current()
	consensusOptions := consensus.Options{
		time.Second / 1,
		filepath.Join(u.HomeDir, ".pangolin", "bolt.db"),
		0755,
	}

	// Main loop status channel
	program := make(chan int)

	// Kick off gossip loop
	network := net.Bootstrap()
	go consensus.Run(program, network, consensusOptions)

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
