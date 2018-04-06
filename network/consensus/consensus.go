package consensus

import (
	"../net"
	"../utils"
	"time"
	"os"
)

type Options struct {
	Interval time.Duration
	Path     string
	Mode     os.FileMode
}

func Run(p *utils.Program, network net.Network, o Options) {
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
