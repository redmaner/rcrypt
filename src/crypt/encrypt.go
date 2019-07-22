package crypt

import (
	chacha "golang.org/x/crypto/chacha20poly1305"
)

// Encrypt is a method that encrypts data with password, using Coffin.Algorithm
func (c *Coffin) Encrypt(data []byte, password []byte) ([]byte, error) {

	// Switch on Coffin.Mode, and select the appropriate encryption algorithm
	switch c.Mode {
	case CryptCHACHA20:
		return c.encryptCHACHA20(data, password)
	default:
		return c.encryptCHACHA20(data, password)
	}
}

// encryptCHACHA20Poly1305 is a function that encrypts data with password using the chacha20-poly1305 encryption algorithm
func (c *Coffin) encryptCHACHA20(data []byte, password []byte) ([]byte, error) {

	// Make a 256bit key from password
	key := makeKey(password)

	// Create a new block
	aead, err := chacha.NewX(key)
	if err != nil {
		return emptyByte, err
	}

	// Generate a nonce if specified by Coffin.Options
	nonce, err := makeNonce(chacha.NonceSizeX, false)
	if err != nil {
		return emptyByte, err
	}
	if c.Opts.WithNonce {
		nonce, err = makeNonce(chacha.NonceSizeX, true)
		if err != nil {
			return emptyByte, err
		}
		c.Opts.Nonce = nonce
	}

	// Seal data
	ciphertext := aead.Seal(nil, nonce, data, nil)

	// Return the data
	return ciphertext, nil
}
