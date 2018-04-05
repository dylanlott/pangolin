package net

import (
	"time"
	"fmt"
	"math/rand"
)

const NetworkSize = 4

type Network [NetworkSize]*NetNode

type NetNode struct {
	id       int
	peers []*NetNode
	// state State
}

func (n *NetNode) RandomGossip(t time.Time) {
	fmt.Println("n:", n)
	r := rand.Intn(len(n.peers))
	peer := n.peers[r]

	fmt.Println("contacting node:", peer)
	// peer.sync()
}

func Bootstrap() Network {
	network := Network{}

	// Generate nodes
	for i, _ := range network {
		network[i] = new(NetNode)
		*network[i] = NetNode{id: i}
	}

	// Populate nodes' peers
	for i, node := range network {
		node.peers = make([]*NetNode, 0)
		for j,peer := range network {
			if i != j {
				network[i].peers = append(node.peers, peer)
			}
		}
	}

	return network
}
