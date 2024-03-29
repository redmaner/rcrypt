package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"

	chacha "golang.org/x/crypto/chacha20poly1305"
)

// Encrypt is a method that encrypts data using Coffin.Algorithm and Coffin.Opts
func (c *Coffin) Encrypt(data []byte) ([]byte, error) {

	// Switch on Coffin.Mode, and select the appropriate encryption algorithm
	switch c.Mode {
	case CryptCHACHA20:
		return c.encryptCHACHA20(data)
	case CryptAES256:
		return c.encryptAES256(data)
	case CryptRSA:
		return c.encryptRSA(data)
	default:
		return c.encryptCHACHA20(data)
	}
}

// encryptCHACHA20Poly1305 is a function that encrypts data with password using the chacha20-poly1305 encryption algorithm
func (c *Coffin) encryptCHACHA20(data []byte) ([]byte, error) {

	// If password is not supplied, return error
	if len(c.Opts.Password) == 0 {
		return emptyByte, ErrNoPassword
	}

	// Make a 256bit key from password
	key := makeKey(c.Opts.Password)

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

// encryptAES256 is a function that encrypts data with password using the AES256-GCM encryption algorithm
func (c *Coffin) encryptAES256(data []byte) ([]byte, error) {

	// If password is not supplied, return error
	if len(c.Opts.Password) == 0 {
		return emptyByte, ErrNoPassword
	}

	// Make a 256bit key from password
	key := makeKey(c.Opts.Password)

	// Create a new block
	block, err := aes.NewCipher(key)
	if err != nil {
		return emptyByte, err
	}

	// Generate a nonce if specified by Coffin.Options
	nonce, err := makeNonce(12, false)
	if err != nil {
		return emptyByte, err
	}
	if c.Opts.WithNonce {
		nonce, err = makeNonce(12, true)
		if err != nil {
			return emptyByte, err
		}
		c.Opts.Nonce = nonce
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return emptyByte, err
	}

	// Seal data
	ciphertext := aead.Seal(nil, nonce, data, nil)

	// Return the data
	return ciphertext, nil
}

// encryptRSA is a function that encrypts data with public key using the RSA encryption algorithm
func (c *Coffin) encryptRSA(data []byte) ([]byte, error) {

	// If public key is not supplied, return error
	if len(c.Opts.PubKey) == 0 {
		return emptyByte, ErrNoPubKey
	}

	// Unmarshall the RSA public key
	rsaKey, err := UnmarshalPublicKey(bytes.NewBuffer(c.Opts.PubKey))
	if err != nil {
		return emptyByte, err
	}

	// Encrypt the data
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, rsaKey, data, nil)
	if err != nil {
		return emptyByte, err
	}

	// Return the data
	return ciphertext, nil
}
