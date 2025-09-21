package p2p

// handshakerfunc is a function that performs a handshake on a connection
type HandshakerFunc func(Peer) error

func NOPHandshaker(Peer) error {
	return nil
}