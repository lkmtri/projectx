package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/lkmtri/projectx/crypto"
	"github.com/lkmtri/projectx/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	txn := Transaction{}

	return NewBlock(header, []Transaction{txn}, WithHasher(BlockHasher{}))
}

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	block := randomBlock(height)
	privateKey := crypto.GeneratePrivateKey()
	err := block.Sign(privateKey)
	assert.Nil(t, err)
	return block
}

func TestBlockHash(t *testing.T) {
	block := randomBlock(10)
	fmt.Println(block.Hash())
}

func TestBlockVerify(t *testing.T) {
	privateKey := crypto.GeneratePrivateKey()
	block := randomBlock(10)

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
