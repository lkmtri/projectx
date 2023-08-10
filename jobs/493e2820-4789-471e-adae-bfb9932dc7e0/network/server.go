package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transport []Transport
}

type Server struct {
	ServerOpts
	trCh   chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{ServerOpts: opts, trCh: make(chan RPC), quitCh: make(chan struct{}, 1)}
}

func (s *Server) Start() {
	s.initTransports()
	t := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.trCh:
			fmt.Printf("%s %s\n", rpc.From, rpc.Payload)
		case <-s.quitCh:
			break free
		case <-t.C:
		}
	}
}

func (s *Server) initTransports() {
	for _, tr := range s.Transport {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.trCh <- rpc
			}
		}(tr)
	}
}
