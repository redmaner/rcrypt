package crypt

import "errors"

var (
	//ErrNoPassword is returned if a password is not supplied
	ErrNoPassword = errors.New("password not supplied")

	//ErrNoNonce is returned when the option WithNonce is enabled but the nonce is not supplied.
	//This error is only returned when decrypting data
	ErrNoNonce = errors.New("nonce not supplied")

	//ErrNoPubKey is returned when no public key was supplied
	ErrNoPubKey = errors.New("public key not supplied")

	//ErrNoPrivKey is returned when no private key was supplied
	ErrNoPrivKey = errors.New("private key not supplied")
)
