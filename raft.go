package raft

import (
	"sync"
	"time"
)

type Role int
const (
	Follower Role = iota
	Candidate
	Leader
)

type RaftNode struct {
	mu sync.Mutex
	ID string
	CurrentTerm int
	VotedFor string
	Role Role
	HeartbeatInterval time.Duration
}

func NewRaftNode(id string) *RaftNode {
	return &RaftNode{
		ID: id,
		CurrentTerm: 0,
		Role: Follower,
		HeartbeatInterval: 150 * time.Millisecond,
	}
}
