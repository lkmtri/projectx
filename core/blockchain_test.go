package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *BlockChain {
	bc, err := NewBlockChain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

func TestNewBlockChain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(uint32(0)))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenBlocks := 1000
	for i := 0; i < lenBlocks; i += 1 {
		err := bc.AddBlock(randomBlockWithSignature(t, uint32(i+1)))
		assert.Nil(t, err)
	}

	assert.Equal(t, bc.Height(), uint32(lenBlocks))

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, uint32(lenBlocks))))

	assert.NotNil(t, bc.AddBlock(randomBlock(uint32(lenBlocks+1))))
	assert.Equal(t, bc.Height(), uint32(lenBlocks))
}
