package consensus

import (
	"time"
	"../net"
)

func Run(c chan int, node *net.NetNode, interval time.Duration) {
	// Loop forever every `interval`
	ticker := time.NewTicker(interval)
	for t := range ticker.C {
		node.RandomGossip(t)
	}
}
