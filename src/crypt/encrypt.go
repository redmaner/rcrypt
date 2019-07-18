package crypt

import (
	chacha "golang.org/x/crypto/chacha20poly1305"
)

func (c *Coffin) Encrypt(data []byte, password []byte) ([]byte, error) {
	switch c.Mode {
	case CryptCHACHA20Poly1305:
		return c.encryptCHACHA20Poly1305(data, password)
	default:
		return c.encryptCHACHA20Poly1305(data, password)
	}
}

func (c *Coffin) encryptCHACHA20Poly1305(data []byte, password []byte) ([]byte, error) {

	key := makeKey(password)
	aead, err := chacha.NewX(key)
	if err != nil {
		return []byte{}, err
	}

	nonce, err := makeNonce(chacha.NonceSizeX)
	if err != nil {
		return []byte{}, err
	}

	ciphertext := aead.Seal(nil, nonce, data, nil)

	return ciphertext, nil
}
