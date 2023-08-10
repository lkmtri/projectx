package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *BlockChain
}

var (
	errBlockAlreadyExisted  = fmt.Errorf("block already exited")
	errBlockTooHigh         = fmt.Errorf("block too high")
	errInvalidPrevBlockHash = fmt.Errorf("invalid prev block hash")
)

func NewBlockValidator(bc *BlockChain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("%w block %d with hash %s", errBlockAlreadyExisted, b.Height, b.Hash())
	}

	if v.bc.Height()+1 < b.Height {
		return fmt.Errorf("%w %s", errBlockTooHigh, b.Hash())
	}

	if prevHash := (BlockHasher{}).Hash(v.bc.GetTailHeader()); prevHash != b.PrevBlockHash {
		return fmt.Errorf("%w %s", errInvalidPrevBlockHash, b.Hash())
	}

	if err := b.Verify(); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	return nil
}
