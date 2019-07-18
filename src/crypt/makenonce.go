package crypt

import (
	"crypto/rand"
	"io"
)

func makeNonce(nonceSize int) ([]byte, error) {
	nonce := make([]byte, nonceSize)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}

	return nonce, nil
}
