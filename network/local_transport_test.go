package network

import (
	"blockchain/test"
	"testing"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	// assert.Equal(t, tra.peers[trb.Addr()], trb)
	// assert.Equal(t, trb.peers[tra.Addr()], tra)

}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	test.AssertNil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <-trb.Consume()
	test.AsserEqual(t, rpc.Payload, msg)
	test.AsserEqual(t, rpc.From, tra.Addr())
}
