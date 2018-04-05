package consensus

import (
	"time"
	"../net"
	"github.com/boltdb/bolt"
	"log"
	"os"
	"path/filepath"
)

type Options struct {
	Interval time.Duration
	Path     string
	Mode os.FileMode
}

var Db *bolt.DB

func Run(c chan int, node *net.NetNode, o Options) {
	ensureDirectory(filepath.Dir(o.Path), o.Mode)
	db, err := bolt.Open(o.Path, o.Mode, nil)

	if err != nil {
		log.Fatal("Persistence error: ", err)
		c <-1
		close(c)
	}

	Db = db

	// Loop forever every `interval`
	ticker := time.NewTicker(o.Interval)
	for t := range ticker.C {
		node.RandomGossip(t)
	}
}
