package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPeer is a struct that represents a remote node in the network via TCP
type TCPeer struct {
	// conn is the underlying TCP connection
	conn     net.Conn
	// if we dial and retrieve the connection, outbound is true
	// if we accept a connection, outbound is false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPeer {
	return &TCPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOptions struct {
	ListenAddr     string
	HandshakerFunc HandshakerFunc
	Decoder	       Decoder
}

type TCPTransport struct {
	TCPTransportOptions
	listener         net.Listener
	mu               sync.RWMutex
	peers            map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		panic(err)
	}

	go t.StartAcceptLoop()

	return nil
}

func (t *TCPTransport) StartAcceptLoop() {

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}

		fmt.Printf("new incoming connection: %+v",conn)

		go t.handleconn(conn)
	}

}

type Temp struct {}

func (t *TCPTransport) handleconn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakerFunc(peer); err != nil {
		conn.Close()
		fmt.Println("Error during handshake:", err)
		return
	 }

	msg := &Message{}

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Println("Error decoding message:", err)
			continue
		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("Received message: %+v\n", msg)
	}
}