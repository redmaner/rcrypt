package crypt

// Algorithm holds the encryption / decryption mode
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

	// Password is the password used for encryption and decryption. This only applies
	// to AES256-GCM and chacha20-poly1305
	Password []byte

	// WithNonce enables the use of a unique nonce. This option can be used with AES and CHACHA20
	// Use this option with encryption. The generated nonce can be accesed with Nonce or the GetNonce() function
	WithNonce bool

	// Nonce is the nonce generated during encryption, and should be supplied for decryption
	Nonce []byte

	// PrivKey is the private key in PEM encoded bytes. Used for decryption with RSA algorithm
	PrivKey []byte

	// PubKey is the public key in PEM encoded bytes. Used for encryption with RSA algorithm
	PubKey []byte
}
