package crypt

import (
	"fmt"
	"testing"
)

func TestCHACHA20(t *testing.T) {

	testData := "This is some test data, written in a string"

	// Make a new Coffin
	cof := NewCoffin(CryptCHACHA20)

	// Set a password
	cof.Opts.Password = []byte("This is a test password, don't use in production")
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
