package net

import (
	"../utils"
	"time"
	"fmt"
	"math/rand"
	"github.com/boltdb/bolt"
)

const NetworkSize = 4

type Network [NetworkSize]*NetNode

type NetNode struct {
	Id    int
	peers []*NetNode
	State State
}

func (n *NetNode) RandomGossip(t time.Time) {
	// fmt.Println("n:", n)
	r := rand.Intn(len(n.peers))
	peer := n.peers[r]

	fmt.Printf("node %d; contacting node: %d\n", n.Id, peer.Id)
	peer.sync(n)
}

func (n *NetNode) sync(peer *NetNode) {
	selfDiff, peerDiff := n.State.diff(&peer.State)
	n.State.apply(&selfDiff)
	peer.State.apply(&peerDiff)
}

func Bootstrap(p *utils.Program, db *bolt.DB) (network Network) {
	network = Network{}

	// Generate nodes
	for i, _ := range network {
		network[i] = new(NetNode)
		*network[i] = NetNode{Id: i}
	}

	// Populate nodes' peers
	for i, node := range network {
		node.State = State{db, []byte(fmt.Sprintf("node-%d", i))}
		node.peers = make([]*NetNode, 0)
		for j, peer := range network {
			if i != j {
				network[i].peers = append(node.peers, peer)
			}
		}
	}

	bucketErrors := EnsureBuckets(db, network)
	for _, err := range bucketErrors {
		p.ErrCheck( err)
	}

	return
}

func EnsureBuckets(db *bolt.DB, network Network) (errors []error) {
	errors = make([]error, 0)

	for _, node := range network {
		db.Update(func(tx *bolt.Tx) (err error) {
			_, err = tx.CreateBucketIfNotExists([]byte(fmt.Sprintf("node-%d", node.Id)))
			errors = append(errors, err)
			return
		})
	}

	return
}

