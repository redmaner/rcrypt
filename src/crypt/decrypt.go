package crypt

import chacha "golang.org/x/crypto/chacha20poly1305"

// Decrypt is a method that decrypts data using Coffin.Algorithm and Coffin.Opts
func (c *Coffin) Decrypt(data []byte) ([]byte, error) {
	switch c.Mode {
	case CryptCHACHA20:
		return c.decryptCHACHA20(data)
	default:
		return c.decryptCHACHA20(data)
	}
}

func (c *Coffin) decryptCHACHA20(data []byte) ([]byte, error) {

	if len(c.Opts.Password) == 0 {
		return emptyByte, ErrNoPassword
	}

	key := makeKey(c.Opts.Password)
	aead, err := chacha.NewX(key)
	if err != nil {
		return []byte{}, err
	}

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

	plaintext, err := aead.Open(nil, nonce, data, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
