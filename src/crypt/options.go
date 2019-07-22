package crypt

// Algoritm holds the encryption / decryption mode
type Algorithm int

const (

	// CryptCHACHA20 uses the chacha20-poly1305 symmetric encryption algorithm
	CryptCHACHA20 Algorithm = 1

	// CryptAES256 uses the AES-GCM 256bit symmetric encryption algoritm
	CryptAES256 Algorithm = 2

	// CryptRSA uses the RSA asymetric encryption algorithm
	CryptRSA Algorithm = 3
)

var emptyByte []byte

// Options specifies options
type Options struct {

	// WithNonce enables the use of a unique nonce
	// This option can be used with AES and CHACHA20
	WithNonce bool
	Nonce     []byte
}
