package p2p

import "net"

// Message represents a generic message structure used in P2P communication
type Message struct {
	From    net.Addr
	Payload []byte
}