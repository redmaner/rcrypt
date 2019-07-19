package crypt

// Algoritm holds the encryption / decryption mode
type Algoritm int

const (

	// CryptCHACHA20 uses the chacha20-poly1305 symmetric encryption algorithm
	CryptCHACHA20 Algoritm = 1

	// CryptAES256 uses the AES-GCM 256bit symmetric encryption algoritm
	CryptAES256 Algoritm = 2

	// CryptRSA uses the RSA asymetric encryption algorithm
	CryptRSA Algoritm = 3
)

var emptyByte []byte

// Options specifies options
type Options struct {

	// WithNonce enables the use of a unique nonce
	// This option can be used with AES and CHACHA20
	WithNonce bool
	nonce     []byte
}
