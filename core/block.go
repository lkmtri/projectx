package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/lkmtri/projectx/crypto"
	"github.com/lkmtri/projectx/types"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Timestamp     int64
	Height        uint32
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// to cache the hash of Header
	hash   types.Hash
	hasher Hasher[*Block]
}

type BlockOpts = func(*Block)

func WithHasher(hasher Hasher[*Block]) BlockOpts {
	return func(b *Block) {
		b.hasher = hasher
	}
}

func NewBlock(h *Header, txx []Transaction, opts ...BlockOpts) *Block {
	block := &Block{
		Header:       h,
		Transactions: txx,
	}

	for _, opt := range opts {
		opt(block)
	}

	return block
}

func (b *Block) Sign(privateKey crypto.PrivateKey) error {
	sig, err := privateKey.Sign(b.hashBytes())
	if err != nil {
		return err
	}

	b.Validator = privateKey.PublicKey()
	b.Signature = sig
	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.hashBytes()) {
		return fmt.Errorf("block has invalid signature")
	}

	return nil
}

func (b *Block) Encode(w io.Writer, encoder Encoder[*Block]) error {
	return encoder.Encode(w, b)
}

func (b *Block) Decode(r io.Reader, decoder Decoder[*Block]) error {
	return decoder.Decode(r, b)
}

func (b *Block) Hash() types.Hash {
	if b.hash.IsZero() {
		b.hash = b.hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) hashBytes() []byte {
	hash := b.hasher.Hash(b)
	return hash.Bytes()
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(*b.Header)
	return buf.Bytes()
}
