package crypt

type CryptMode int

const (
	CryptCHACHA20Poly1305 CryptMode = 1
	CryptAES256           CryptMode = 2
	CryptRSA              CryptMode = 3
)
