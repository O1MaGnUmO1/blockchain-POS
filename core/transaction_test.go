package core

import (
	"blockchain/crypto"
	"blockchain/test"
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
	tx.PublicKey = otherPrivkey.PublicKey()

	test.AssertNotNil(t, tx.Verify())
}
