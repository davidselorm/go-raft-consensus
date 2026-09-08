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

func (r Role) String() string {
	switch r {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	default:
		return "Unknown"
	}
}

type LogEntry struct {
	Index uint64
	Term  uint64
	Data  []byte
}

type ApplyMsg struct {
	CommandValid bool
	Command      []byte
	CommandIndex uint64
	SnapshotValid bool
	Snapshot      []byte
	SnapshotTerm  uint64
	SnapshotIndex uint64
}

type RaftNode struct {
	mu        sync.Mutex
	peers     []*RaftNode
	id        int
	currentTerm uint64
	votedFor  int
	log       []LogEntry
	commitIdx uint64
	lastApplied uint64
	role      Role
	heartbeat time.Duration
	applyCh   chan ApplyMsg
	stopCh    chan struct{}
}

func NewRaftNode(id int, peersCount int, applyCh chan ApplyMsg) *RaftNode {
	rn := &RaftNode{
		id:          id,
		currentTerm: 0,
		votedFor:    -1,
		log:         make([]LogEntry, 0),
		commitIdx:   0,
		lastApplied: 0,
		role:        Follower,
		heartbeat:   50 * time.Millisecond,
		applyCh:     applyCh,
		stopCh:      make(chan struct{}),
	}
	// Initial dummy entry at index 0
	rn.log = append(rn.log, LogEntry{Index: 0, Term: 0, Data: nil})
	return rn
}

func (rn *RaftNode) GetState() (uint64, bool) {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	return rn.currentTerm, rn.role == Leader
}

func (rn *RaftNode) Stop() {
	close(rn.stopCh)
}
