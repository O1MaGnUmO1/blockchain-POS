package core

import (
	"blockchain/crypto"
	"blockchain/types"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Timestamp     uint64
	Height        uint32
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
	// Cashed version of the Header hash
	hash types.Hash
}

func NewBlock(h *Header, txs []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txs,
	}
}

func (b *Block) Sign(privateKey crypto.PrivateKey) *crypto.Signature {
	sig, err := privateKey.Sign(b.HeaderData())
	if err != nil {
		panic(err)
	}
	b.Validator = privateKey.PublicKey()
	b.Signature = sig
	return nil
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}
	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("block has invalid signature")
	}
	return nil
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}
	return b.hash
}

// func (b *Block) Hash() types.Hash {
// 	buf := &bytes.Buffer{}
// 	b.Header.EncodeBinary(buf)

// 	if b.hash.IsZero() {
// 		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
// 	}
// 	return b.hash
// }
