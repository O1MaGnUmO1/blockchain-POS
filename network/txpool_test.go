package network

import (
	"blockchain/core"
	"blockchain/test"
	"math/rand"
	"strconv"
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

func TestSortTransactions(t *testing.T) {
	p := NewTxPool()

	txLen := 1000
	for i := 0; i < txLen; i++ {
		tx := core.NewTransaction([]byte(strconv.FormatInt(int64(i), 10)))
		tx.SetFirstSeen(int64(i * rand.Intn(10000)))
		test.AssertNil(t, p.Add(tx))
	}
	test.AsserEqual(t, txLen, p.Len())

	txx := p.Transactions()
	for i := 0; i < len(txx)-1; i++ {
		test.AssertTrue(t, txx[i].FirstSeen() < txx[i+1].FirstSeen())
	}
}
