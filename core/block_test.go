package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/lkmtri/projectx/crypto"
	"github.com/lkmtri/projectx/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(t *testing.T, height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	return NewBlock(header, []Transaction{}, WithHasher(BlockHasher{}))
}

func randomTxWithSignature(t *testing.T) Transaction {
	txn := Transaction{}
	assert.Nil(t, txn.Sign(crypto.GeneratePrivateKey()))
	return txn
}

func signBlock(t *testing.T, block *Block) {
	privateKey := crypto.GeneratePrivateKey()
	err := block.Sign(privateKey)
	assert.Nil(t, err)
}

func TestBlockHash(t *testing.T) {
	block := randomBlock(t, 10)
	fmt.Println(block.Hash())
}

func TestBlockVerify(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	block := randomBlock(t, 10)

	assert.NotNil(t, block.Verify())

	assert.Nil(t, block.Sign(privateKey))
	assert.NotNil(t, block.Signature)

	otherPrivateKey := crypto.GeneratePrivateKey()
	otherPublicKey := otherPrivateKey.PublicKey()

	assert.Nil(t, block.Verify())
	block.Validator = otherPublicKey
	assert.NotNil(t, block.Verify())

	block.Validator = privateKey.PublicKey()
	assert.Nil(t, block.Verify())
	block.Version = 2
	assert.NotNil(t, block.Verify())
}
