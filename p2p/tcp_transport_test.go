package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcpTransport(t *testing.T) {
	// Use :0 so the OS picks an available port, avoiding conflicts during tests
	listenaddr := ":0"
	tr := NewTCPTransport(TCPTransportOptions{
		ListenAddr:     listenaddr,
		HandshakerFunc: NOPHandshaker,
		Decoder:        GOBDecoder{},
	})

	assert.NotNil(t, tr)
	assert.Equal(t, listenaddr, tr.ListenAddr)

	// Start listening
	err := tr.ListenAndAccept()
	assert.NoError(t, err)
}
