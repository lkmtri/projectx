package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/lkmtri/projectx/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{r, s}, nil
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (k PublicKey) Address() types.Address {
	b := sha256.Sum256(elliptic.MarshalCompressed(elliptic.P256(), k.key.X, k.key.Y))
	return types.AddressFromBytes(b[len(b)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (s Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.key, data, s.r, s.s)
}
