package raft

import (
	"testing"
)

func TestRaftNodeInitialization(t *testing.T) {
	applyCh := make(chan ApplyMsg, 16)
	node := NewRaftNode(0, 3, applyCh)
	defer node.Stop()

	term, isLeader := node.GetState()
	if term != 0 {
		t.Fatalf("Expected initial term 0, got %d", term)
	}
	if isLeader {
		t.Fatalf("Initial state must be Follower, not Leader")
	}
}

func TestRequestVoteTermPrecedence(t *testing.T) {
	applyCh := make(chan ApplyMsg, 16)
	node := NewRaftNode(0, 3, applyCh)
	defer node.Stop()

	args := &RequestVoteArgs{
		Term:         2,
		CandidateId:  1,
		LastLogIndex: 0,
		LastLogTerm:  0,
	}
	reply := &RequestVoteReply{}
	node.RequestVote(args, reply)

	if !reply.VoteGranted {
		t.Fatalf("Expected vote to be granted for higher term")
	}
	if reply.Term != 2 {
		t.Fatalf("Expected node to adopt term 2")
	}
}
