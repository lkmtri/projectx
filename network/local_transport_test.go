package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalTransport(t *testing.T) {
	tra := NewLocalTransport("a")
	trb := NewLocalTransport("b")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("a")
	trb := NewLocalTransport("b")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Hello World")
	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.Addr())
}
