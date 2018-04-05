package net

import (
	"time"
	"fmt"
	"math/rand"
)

const NetworkSize = 4

type Network [NetworkSize]*NetNode

type NetNode struct {
	Id    int
	peers []*NetNode
	state State
}

func (n *NetNode) RandomGossip(t time.Time) {
	// fmt.Println("n:", n)
	r := rand.Intn(len(n.peers))
	peer := n.peers[r]

	fmt.Printf("node %d; contacting node: %d\n", n.Id, peer.Id)
	peer.sync(n)
}

func (n *NetNode) sync(peer *NetNode) {
	selfDiff, peerDiff := n.state.diff(&peer.state)
	n.state.apply(&selfDiff)
	peer.state.apply(&peerDiff)
}

func Bootstrap() Network {
	network := Network{}

	// Generate nodes
	for i, _ := range network {
		network[i] = new(NetNode)
		*network[i] = NetNode{Id: i}
	}

	// Populate nodes' peers
	for i, node := range network {
		node.peers = make([]*NetNode, 0)
		for j, peer := range network {
			if i != j {
				network[i].peers = append(node.peers, peer)
			}
		}
	}

	return network
}
