package core

import (
	"blockchain/crypto"
	"blockchain/test"
	"bytes"
	"testing"
)

func TestSignTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("data"),
	}
	test.AssertNil(t, tx.Sign(privateKey))
	test.AssertNotNil(t, tx.Signature)

}

func TestVerifyTransaction(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("Data"),
	}
	test.AssertNil(t, tx.Sign(privateKey))
	test.AssertNil(t, tx.Verify())

	otherPrivkey := crypto.GeneratePrivateKey()
	tx.From = otherPrivkey.PublicKey()

	test.AssertNotNil(t, tx.Verify())
}

func TestTxEncodeDecode(t *testing.T) {
	tx := randomTxWithSignature(t)
	buf := &bytes.Buffer{}
	// gob.Register(elliptic.P256())
	test.AssertNil(t, tx.Encode(NewGobTxEncoder(buf)))

	txDecoded := new(Transaction)
	test.AssertNil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
	test.AsserEqual(t, tx, txDecoded)
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privateKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("hello"),
	}
	test.AssertNil(t, tx.Sign(privateKey))
	return tx
}
