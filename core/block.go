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

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	_ = enc.Encode(h)
	return buf.Bytes()
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// to cache the hash of Header
	hash   types.Hash
	hasher Hasher[*Header]
}

type BlockOpts = func(*Block)

var (
	errBlockNoSignature = fmt.Errorf("block has no signature")
	errInvalidSignature = fmt.Errorf("block has invalid signature")
)

func WithHasher(hasher Hasher[*Header]) BlockOpts {
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

func (b *Block) AddTransaction(txn Transaction) {
	b.Transactions = append(b.Transactions, txn)
}

func (b *Block) Sign(privateKey crypto.PrivateKey) error {
	if b.Signature != nil {
		return nil
	}

	sig, err := privateKey.Sign(b.Bytes())
	if err != nil {
		return err
	}

	b.Validator = privateKey.PublicKey()
	b.Signature = sig
	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return errBlockNoSignature
	}

	if !b.Signature.Verify(b.Validator, b.Bytes()) {
		return errInvalidSignature
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
		b.hash = b.hasher.Hash(b.Header)
	}

	return b.hash
}

func (b *Block) Bytes() []byte {
	hash := b.hasher.Hash(b.Header)
	return hash.Bytes()
}
