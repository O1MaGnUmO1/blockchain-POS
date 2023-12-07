package core

import (
	"blockchain/crypto"
	"blockchain/test"
	"blockchain/types"
	"testing"
	"time"
)

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}
	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privkey := crypto.GeneratePrivateKey()
	bc := randomBlock(height, prevBlockHash)
	test.AssertNil(t, bc.Sign(privkey))

	return bc
}

func TestSignBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	test.AssertNil(t, b.Sign(privateKey))
	test.AssertNotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	test.AssertNil(t, b.Sign(privateKey))
	test.AssertNil(t, b.Verify())
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(uint32(height) - 1)
	test.AssertNil(t, err)
	return BlockHasher{}.Hash(prevHeader)
}
