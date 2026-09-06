package raft

import (
	"math/rand"
	"time"
)

func (rn *RaftNode) RandomizedTimeout() time.Duration {
	return time.Duration(150+rand.Intn(150)) * time.Millisecond
}

func (rn *RaftNode) StartElection() {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.CurrentTerm++
	rn.Role = Candidate
	rn.VotedFor = rn.ID
}
