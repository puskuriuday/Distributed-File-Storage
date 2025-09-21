package main

import (
	
	"log"

	"github.com/puskuriuday/Distributed-File-Storage/p2p"
)

func main() {
    opts := p2p.TCPTransportOptions{
		ListenAddr:     ":4000",
		HandshakerFunc: p2p.NOPHandshaker,
		Decoder:        p2p.GOBDecoder{},
	}

	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}