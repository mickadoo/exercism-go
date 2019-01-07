package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a new random private key from given limit
func PrivateKey(limit *big.Int) *big.Int {
	min := big.NewInt(2)
	max := big.NewInt(0)
	max = max.Add(max, limit)
	max = max.Sub(max, min)
	randNum, err := rand.Int(rand.Reader, max)

	if err != nil {
		panic(err)
	}

	return randNum.Add(randNum, min)
}

// PublicKey generates a new public key based on a provided private key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	gBig := big.NewInt(g)

	return gBig.Exp(gBig, private, p)
}

// NewPair creates a new public and private key pair
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)

	return private, PublicKey(private, p, g)
}

// SecretKey gets a secret from given big int
func SecretKey(private1, public2, p *big.Int) *big.Int {
	a := big.NewInt(0)

	return a.Exp(public2, private1, p)
}
