package core

import (
	"blockchain/crypto"
	"blockchain/test"
	"blockchain/types"
	"testing"
	"time"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}
	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	privkey := crypto.GeneratePrivateKey()
	bc := randomBlock(height)
	test.AssertNil(t, bc.Sign(privkey))

	return bc
}

func TestSignBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	test.AssertNil(t, b.Sign(privateKey))
	test.AssertNotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)
	test.AssertNil(t, b.Sign(privateKey))
	test.AssertNil(t, b.Verify())
}
