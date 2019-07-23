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

	fmt.Printf("Tests for chacha20 passed succesfully")
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

	fmt.Printf("Tests for aes256 passed succesfully")
}
