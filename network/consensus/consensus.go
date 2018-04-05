package consensus

import (
	"../net"
	"../utils"
	"time"
	"github.com/boltdb/bolt"
	"log"
	"os"
	"path/filepath"
)

type Options struct {
	Interval time.Duration
	Path     string
	Mode     os.FileMode
}

var Db *bolt.DB

func Run(c chan int, network net.Network, o Options) {
	err := utils.EnsureDirectory(filepath.Dir(o.Path), o.Mode | os.ModeDir)
	errCheck(c, err)
	db, err := bolt.Open(o.Path, o.Mode, nil)
	errCheck(c, err)

	Db = db

	for _, node := range network {
		go (func(node *net.NetNode) {
			// Loop forever every `interval`
			ticker := time.NewTicker(o.Interval)
			for t := range ticker.C {
				node.RandomGossip(t)
			}
		})(node)
	}
}

func errCheck(c chan int, err error) {
	if err != nil {
		log.Fatal("Persistence error: ", err)
		c <- 1
		close(c)
	}
}
