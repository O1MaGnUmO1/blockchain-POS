package core

import (
	"blockchain/types"
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     989654,
	}
	buf := bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(&buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(&buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989654,
		},
		Transactions: nil,
	}
	fmt.Printf("%v", b)
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989654,
		},
		Transactions: nil,
	}

	h := b.Hash()
	assert.False(t, h.IsZero())
}
