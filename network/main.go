package main

import (
	"./utils"
	"./net"
	"./consensus"
	"os"
	"os/signal"
	"os/user"
	"fmt"
	"time"
	"path/filepath"
	"github.com/boltdb/bolt"
)

func main() {
	// Initialize
	program, db, consensusOptions := initialize()


	// Kick off gossip loop
	network := net.Bootstrap(&program, db)
	go consensus.Run(&program, network, consensusOptions)

	// Block until `program` is closed or SIGINT
	go handleSignals(&program)
	os.Exit(block(&program, 0))
}

func initialize() (p utils.Program, db *bolt.DB, o consensus.Options) {
	// Main loop status channel
	p = utils.Program{make(chan int)}

	u, _ := user.Current()
	o = consensus.Options{
		time.Second / 1,
		filepath.Join(u.HomeDir, ".pangolin", "bolt.db"),
		0755,
	}

	err := utils.EnsureDirectory(filepath.Dir(o.Path), o.Mode|os.ModeDir)
	p.ErrCheck(err)

	db, err = bolt.Open(o.Path, o.Mode, nil)
	p.ErrCheck(err)

	return p, db, o
}

func block(p *utils.Program, status int) int {
	_status, ok := <-p.C
	if !ok {
		fmt.Println("last status:", _status)
		return status
	}

	return block(p, _status)
}

func handleSignals(p *utils.Program) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for range interrupt {
		fmt.Printf("Stopping... ")
		close(p.C)
	}
}
