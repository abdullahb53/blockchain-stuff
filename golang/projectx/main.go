package main

import (
	"time"

	"github.com/abdullahb53/blockchain-stuff/projectx/network"
)

// Server
// Transport layer => tcp, udp,
// Block
// Transactions
// Keypair

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello world trRemote"))
			time.Sleep(time.Second * 1)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
