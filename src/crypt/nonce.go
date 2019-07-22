package crypt

import (
	"crypto/rand"
	"io"
)

// makeNonce is a function that generates a random nonce
func makeNonce(nonceSize int, fillRand bool) ([]byte, error) {
	nonce := make([]byte, nonceSize)

	if fillRand {
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return []byte{}, err
		}
	}

	return nonce, nil
}

// GetNonce returns the nonce generated for use with AES or CHACHA20
func (c *Coffin) GetNonce() []byte {
	return c.Opts.Nonce
}
