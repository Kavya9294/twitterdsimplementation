package main

import (
	"context"
	"time"

	"github.com/coreos/etcd/raft"
	"go.etcd.io/etcd/raft/raftpb"
)

type Cluster struct {
	peers []raft.Peer{}
}

type Node struct {
	raft.Node
	Store  *raft.MemoryStorage
	ticker *time.Ticker
	Ctx    context.Context
	ID     uint64
	//Address string
	//Port    int
	//Error   error
	//Client   *Raft
	Cluster *Cluster
	//Server   *grpc.Server
	//Listener net.Listener
}

func raftInit() {

	storage := raft.NewMemoryStorage()
	c := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}

	//Starting a 3 node cluster
	//TODO: Start the other 2 nodes

	n := raft.StartNode(c, []raft.Peer{{ID: 0x02}, {ID: 0x03}})

	//peers := []raft.Peer{{ID: 0x01}}

}

//TODO: Add nodes to the cluster

//func (n *Node)AddNodes( ctx context.Context, cc raftpb.ConfChange ){

//err := n.ProposeConfChange(ctx, cc)
//if err !=nil {
////&AddNodeResponse{
////Success:false,
////Error: ErrConfChangeRefused.Error(),
////},nil
//fmt.Print("Error with adding node:", err)
//}
//for _,node := range n.Cluster.Peers() {
//nodes = append(nodes, )
//}

//}

func (n *Node) recvRaftRPC(ctx context.Context, m raftpb.Message) {
	n.Step(ctx, m)
}

func (n *Node) saveToStorage(hardState raftpb.HardState, entries []raftpb.Entry, snapshot raftpb.Snapshot) {
	n.Store.Append(entries)
	if !raft.IsEmptyHardState(hardState) {
		n.Store.SetHardState(hardState)
	}

	if !raft.IsEmptySnap(snapshot) {
		n.Store.ApplySnapshot(snapshot)
	}
}

func (n *Node) send(messages []raftpb.Message) {
	peers := n.Cluster.peers

	for _, m := range messages {
		// Process locally
		if m.To == n.ID {
			n.Step(n.Ctx, m)
			continue
		}

		// If node is an active raft member send the message
		if _, peer := range peers {
			_, err := peer.Client.Send(n.Ctx, &m)
			if err != nil {
				n.ReportUnreachable(peer.ID)
			}
		}
	}
}

func (n *Node) raftMaintain() {
	for {
		select {
		case <-n.Ticker.C:
			n.Tick()
		case rd := <-n.Ready():
			n.saveToStorage(rd.HardState, rd.Entries, rd.Snapshot)
			n.Send(rd.Messages)
			if !raft.IsEmptySnap(rd.Snapshot) {
				processSnapshot(rd.Snapshot)
			}
			for _, entry := range rd.CommittedEntries {
				n.process(entry)
				if entry.Type == raftpb.EntryConfChange {
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)
					n.ApplyConfChange(cc)
				}
			}
			n.Advance()
		case <-n.done:
			return
		}
	}

}

func raftResume() {
	storage := raft.NewMemoryStorage()

	// Recover the in-memory storage from persistent snapshot, state and entries.
	//storage.ApplySnapshot(snapshot)
	//storage.SetHardState(state)
	//storage.Append(entries)

	c := &Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}

	// Restart raft without peer information.
	// Peer information is already included in the storage.
	n := raft.RestartNode(c)
}
