package core

import (
	"testing"

	"github.com/lkmtri/projectx/crypto"
	"github.com/stretchr/testify/assert"
)

func TestTransactionSign(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	msg := []byte("Hello World")
	txn := &Transaction{
		Data: msg,
	}

	err := txn.Sign(privateKey)
	assert.Nil(t, err)
	assert.NotNil(t, txn.Signature)
}

func TestTransactionVerify(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	msg := []byte("Hello World")
	txn := &Transaction{
		Data: msg,
	}

	err := txn.Sign(privateKey)
	assert.Nil(t, err)
	assert.Nil(t, txn.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	// tamper with public key
	txn.Validator = otherPublicKey
	assert.NotNil(t, txn.Verify())

	// tamper with data
	txn.Validator = privateKey.PublicKey()
	assert.Nil(t, txn.Verify())
	txn.Data = []byte("Tampered data")
	assert.NotNil(t, txn.Verify())
}
