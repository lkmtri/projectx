package types

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

type Hash [32]uint8

func (h *Hash) IsZero() bool {
	for i := 0; i < 32; i += 1 {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

func (h *Hash) ToSlice() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i += 1 {
		b[i] = h[i]
	}
	return b
}

func (h *Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("given bytes with length %d should be 32", len(b))
		panic(msg)
	}

	var value [32]uint8
	for i := 0; i < 32; i += 1 {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
