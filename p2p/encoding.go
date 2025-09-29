package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct{}

func (d GOBDecoder) Decode(r io.Reader, M *Message) error {
	return gob.NewDecoder(r).Decode(M)
}

type DefaultDecoding struct{}

func (d DefaultDecoding) Decode(r io.Reader, M *Message) error {
	buff := make([]byte, 1024)
	n, err := r.Read(buff)
	if err != nil {
		return err
	}

	M.Payload = buff[:n]
	fmt.Println("Received bytes:", string(buff[:n]))
	return nil
}
