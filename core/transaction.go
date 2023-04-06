package core

import (
	"fmt"
	"io"

	"github.com/lkmtri/projectx/crypto"
)

type Transaction struct {
	Data      []byte
	Validator crypto.PublicKey
	Signature *crypto.Signature
}

func (t *Transaction) Sign(privateKey crypto.PrivateKey) error {
	sig, err := privateKey.Sign(t.Data)
	if err != nil {
		return err
	}

	t.Signature = sig
	t.Validator = privateKey.PublicKey()
	return nil
}

func (t *Transaction) Verify() error {
	if t.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !t.Signature.Verify(t.Validator, t.Data) {
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
