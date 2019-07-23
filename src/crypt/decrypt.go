package crypt

import (
	"crypto/aes"
	"crypto/cipher"

	chacha "golang.org/x/crypto/chacha20poly1305"
)

// Decrypt is a method that decrypts data using Coffin.Algorithm and Coffin.Opts
func (c *Coffin) Decrypt(data []byte) ([]byte, error) {
	switch c.Mode {
	case CryptCHACHA20:
		return c.decryptCHACHA20(data)
	case CryptAES256:
		return c.decryptAES256(data)
	default:
		return c.decryptCHACHA20(data)
	}
}

func (c *Coffin) decryptCHACHA20(data []byte) ([]byte, error) {

	// When password is not supplied, return error
	if len(c.Opts.Password) == 0 {
		return emptyByte, ErrNoPassword
	}

	// Make 256bit key from password
	key := makeKey(c.Opts.Password)

	// Create a cipher
	aead, err := chacha.NewX(key)
	if err != nil {
		return []byte{}, err
	}

	// Make a nonce
	nonce, err := makeNonce(chacha.NonceSizeX, false)
	if err != nil {
		return []byte{}, err
	}
	if c.Opts.WithNonce {
		if len(c.Opts.Nonce) == 0 {
			return emptyByte, ErrNoNonce
		}
		nonce = c.Opts.Nonce
	}

	// Seal data with cipher
	plaintext, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}

func (c *Coffin) decryptAES256(data []byte) ([]byte, error) {

	// When password is not supplied, return error
	if len(c.Opts.Password) == 0 {
		return emptyByte, ErrNoPassword
	}

	// Make 256bit key from password
	key := makeKey(c.Opts.Password)

	// Create an AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	// Make a nonce
	nonce, err := makeNonce(chacha.NonceSizeX, false)
	if err != nil {
		return []byte{}, err
	}
	if c.Opts.WithNonce {
		if len(c.Opts.Nonce) == 0 {
			return emptyByte, ErrNoNonce
		}
		nonce = c.Opts.Nonce
	}

	// Create GCM
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return emptyByte, err
	}

	// Open data
	plaintext, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
