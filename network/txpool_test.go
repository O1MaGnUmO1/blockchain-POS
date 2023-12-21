package network

import (
	"blockchain/core"
	"blockchain/test"
	"testing"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	test.AsserEqual(t, p.Len(), 0)
}

func TestTxPoolAddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("helooo"))
	test.AssertNil(t, p.Add(tx))
	test.AsserEqual(t, p.Len(), 1)

	_ = core.NewTransaction([]byte("helooo"))
	test.AsserEqual(t, p.Len(), 1)

	p.Flush()
	test.AsserEqual(t, p.Len(), 0)
}
