package crypto

import (
	"blockchain/test"
	"testing"
)

func TestKeypair_Sighn_Verify(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := []byte("Hello world")
	// normal case
	sig, err := privateKey.Sign(msg)
	test.AssertNil(t, err)
	test.AssertNil(t, err)
	test.AssertTrue(t, sig.Verify(publicKey, msg))
	// another private key
	pk := GeneratePrivateKey()
	pb := pk.PublicKey()

	anotherMsg := []byte("Goodbye")
	// verify message of another public key
	test.AssertFalse(t, sig.Verify(pb, msg))
	// verify wrong message of public key
	test.AssertFalse(t, sig.Verify(publicKey, anotherMsg))
}
