package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair_Sighn_Verify(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := []byte("Hello world")
	// normal case
	sig, err := privateKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(publicKey, msg))
	// another private key
	pk := GeneratePrivateKey()
	pb := pk.PublicKey()

	anotherMsg := []byte("Goodbye")
	// verify message of another public key
	assert.False(t, sig.Verify(pb, msg))
	// verify wrong message of public key
	assert.False(t, sig.Verify(publicKey, anotherMsg))
}
