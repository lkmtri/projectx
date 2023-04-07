package core

import (
	"fmt"
	"io"

	"github.com/lkmtri/projectx/crypto"
)

type Transaction struct {
	Data      []byte
	From      crypto.PublicKey // public key of sender
	Signature *crypto.Signature
}

func (t *Transaction) Sign(privateKey crypto.PrivateKey) error {
	if t.Signature != nil {
		return nil
	}

	sig, err := privateKey.Sign(t.Data)
	if err != nil {
		return err
	}

	t.Signature = sig
	t.From = privateKey.PublicKey()
	return nil
}

func (t *Transaction) Verify() error {
	if t.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !t.Signature.Verify(t.From, t.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}

func (t *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (t *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}
