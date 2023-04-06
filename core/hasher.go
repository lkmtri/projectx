package core

import (
	"crypto/sha256"

	"github.com/lkmtri/projectx/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(b *Block) types.Hash {
	hash := sha256.Sum256(b.HeaderData())
	return types.Hash(hash)
}
