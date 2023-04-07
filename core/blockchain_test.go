package core

import (
	"testing"

	"github.com/lkmtri/projectx/types"
	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *BlockChain {
	bc, err := NewBlockChain(randomBlock(t, 0))
	assert.Nil(t, err)
	return bc
}

func randomBlockChainBlock(t *testing.T, bc *BlockChain) *Block {
	block := randomBlock(t, bc.GetTailHeader().Height+1)
	block.PrevBlockHash = (BlockHasher{}).Hash(bc.GetTailHeader())
	block.AddTransaction(randomTxWithSignature(t))
	return block
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
	lenBlocks := 2
	for i := 0; i < lenBlocks; i += 1 {
		block := randomBlockChainBlock(t, bc)
		signBlock(t, block)
		err := bc.AddBlock(block)
		assert.Nil(t, err)
	}

	assert.Equal(t, bc.Height(), uint32(lenBlocks))

	// block existed
	block := randomBlockChainBlock(t, bc)
	block.Height = uint32(lenBlocks)
	signBlock(t, block)
	assert.ErrorIs(t, bc.AddBlock(block), errBlockAlreadyExisted)

	// block too high
	block = randomBlockChainBlock(t, bc)
	block.Height = uint32(lenBlocks + 10)
	signBlock(t, block)
	assert.ErrorIs(t, bc.AddBlock(block), errBlockTooHigh)

	// unsigned block
	block = randomBlockChainBlock(t, bc)
	assert.ErrorIs(t, bc.AddBlock(block), errBlockNoSignature)

	// invalid previos block hash
	block = randomBlockChainBlock(t, bc)
	block.PrevBlockHash = types.RandomHash()
	signBlock(t, block)
	assert.ErrorIs(t, bc.AddBlock(block), errInvalidPrevBlockHash)

	// invalid transaction signature
	block = randomBlockChainBlock(t, bc)
	block.AddTransaction(Transaction{})
	signBlock(t, block)
	assert.NotNil(t, bc.AddBlock(block))

	assert.Equal(t, bc.Height(), uint32(lenBlocks))
}
