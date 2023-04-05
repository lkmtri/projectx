package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifySuccess(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := []byte("Hello World")
	sig, err := privateKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(publicKey, msg))
}

func TestVerifyFailure(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	msg := []byte("Hello World")
	sig, err := privateKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivateKey := GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("other message")))
}
