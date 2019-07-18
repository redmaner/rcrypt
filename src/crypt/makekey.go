package crypt

import "crypto/sha512"

// makeKey is a function that generates a 256bit key from a password using sha256
func makeKey(pass []byte) []byte {
	key := make([]byte, 32)
	hash := sha512.Sum512_256(pass)
	for index, val := range hash {
		key[index] = val
	}
	return key
}
