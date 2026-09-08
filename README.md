# go-raft-consensus

A clean, modular implementation of the Raft distributed consensus state machine in Go.

## Architecture
- **State Transitions**: Follower, Candidate, and Leader states with randomized election backoff timers (150-300ms).
- **Consensus RPCs**: `RequestVote` and `AppendEntries` with term matching and log consistency invariants.
- **Log Replication**: Committed log entry streaming to `applyCh` channel.
