package crypt

import chacha "golang.org/x/crypto/chacha20poly1305"

// Decrypt is a method that decrypts data with password, using Coffin.Algorithm
func (c *Coffin) Decrypt(data []byte, password []byte) ([]byte, error) {
	switch c.Mode {
	case CryptCHACHA20:
		return c.decryptCHACHA20(data, password)
	default:
		return c.decryptCHACHA20(data, password)
	}
}

func (c *Coffin) decryptCHACHA20(data []byte, password []byte) ([]byte, error) {

	key := makeKey(password)
	aead, err := chacha.NewX(key)
	if err != nil {
		return []byte{}, err
	}

	plaintext, err := aead.Open(nil, c.Opts.nonce, data, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
