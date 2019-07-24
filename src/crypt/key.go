package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
)

// NewRSAKeyPair generates a new key pair
func NewRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privkey, &privkey.PublicKey, nil
}

// MarshalPrivateKey is a function to marshal a private key to PEM encoded bytes
func MarshalPrivateKey(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)
	return privBytes
}

// MarshalPublicKey is a function to marshal a public key to PEM encoded bytes
func MarshalPublicKey(pub *rsa.PublicKey) ([]byte, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return emptyByte, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes, nil
}

// UnmarshalPrivateKey to private key
func UnmarshalPrivateKey(keyFile io.Reader) (*rsa.PrivateKey, error) {

	priv, err := ioutil.ReadAll(keyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// UnmarshalPublicKey is a function to convert bytes to *rsa.PublicKey
func UnmarshalPublicKey(keyFile io.Reader) (*rsa.PublicKey, error) {

	pub, err := ioutil.ReadAll(keyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Public key not ok")
	}
	return key, nil
}
