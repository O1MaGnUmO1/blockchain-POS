package core

import (
	"blockchain/test"
	"testing"
)

func NewBlockChainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	test.AssertNil(t, err)
	return bc
}

func TestBlockchain(t *testing.T) {
	bc := NewBlockChainWithGenesis(t)
	test.AsserEqual(t, bc.Height(), uint32(0))
	test.AssertNotNil(t, bc.validator)
}

func TestHasBlock(t *testing.T) {
	bc := NewBlockChainWithGenesis(t)
	test.AssertTrue(t, bc.HasBlock(0))
}

func TestAddBlock(t *testing.T) {
	bc := NewBlockChainWithGenesis(t)
	lenBlock := 1000
	for i := 0; i < lenBlock; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))
		test.AssertNil(t, bc.AddBlock(block))
	}
	test.AsserEqual(t, bc.Height(), uint32(lenBlock))
	test.AsserEqual(t, len(bc.headers), lenBlock+1)

	test.AssertNotNil(t, bc.AddBlock(randomBlock(89)))
}
