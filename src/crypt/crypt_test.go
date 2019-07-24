package crypt

import (
	"fmt"
	"testing"
)

const (
	testData     = "This is some test data, written in a string"
	testPassword = "This is some test data, written in a string"
)

func TestCHACHA20(t *testing.T) {

	// Make a new Coffin
	cof := NewCoffin(CryptCHACHA20)

	// Set a password
	cof.Opts.Password = []byte(testPassword)
	cof.Opts.WithNonce = true

	// Encrypt some test data
	encryptedData, err := cof.Encrypt([]byte(testData))
	if err != nil {
		t.Fail()
	}

	plaintext, err := cof.Decrypt(encryptedData)
	if err != nil {
		t.Fail()
	}

	if testData != string(plaintext) {
		t.Fail()
	}

	fmt.Printf("Tests for chacha20 passed succesfully\n")
}

func TestAES256(t *testing.T) {

	// Make a new Coffin
	cof := NewCoffin(CryptAES256)

	// Set a password
	cof.Opts.Password = []byte(testPassword)
	cof.Opts.WithNonce = true

	// Encrypt some test data
	encryptedData, err := cof.Encrypt([]byte(testData))
	if err != nil {
		t.Fail()
	}

	plaintext, err := cof.Decrypt(encryptedData)
	if err != nil {
		t.Fail()
	}

	if testData != string(plaintext) {
		t.Fail()
	}

	fmt.Printf("Tests for aes256 passed succesfully\n")
}

func TestRSA(t *testing.T) {

	priv, pub, err := NewRSAKeyPair(4096)
	if err != nil {
		t.Fail()
	}

	pubkey, err := MarshalPublicKey(pub)
	if err != nil {
		t.Fail()
	}

	privkey := MarshalPrivateKey(priv)

	cof := NewCoffin(CryptRSA)
	cof.Opts.PrivKey = privkey
	cof.Opts.PubKey = pubkey

	encryptedData, err := cof.Encrypt([]byte(testData))
	if err != nil {
		t.Fail()
	}

	plaintext, err := cof.Decrypt(encryptedData)
	if err != nil {
		t.Fail()
	}

	if string(plaintext) != testData {
		t.Fail()
	}

	fmt.Printf("Tests for RSA passed succesfully\n")
}
