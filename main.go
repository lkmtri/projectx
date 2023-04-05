package main

import (
	"time"

	"github.com/lkmtri/projectx/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello World"))
			time.Sleep(1 * time.Second)
		}
	}()

	serverOpts := network.ServerOpts{
		Transport: []network.Transport{trLocal},
	}

	server := network.NewServer(serverOpts)
	server.Start()
}
