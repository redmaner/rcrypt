package crypt

// Coffin is a type that enables various encryption and decryption methods
type Coffin struct {

	// Mode determines the encryption / decryption method
	Mode Algorithm

	// Opts are options (optional)
	Opts Options
}

// NewCoffin returns a new coffin with Algoritm specified by alg
func NewCoffin(alg Algorithm) *Coffin {
	return &Coffin{
		Mode: alg,
	}
}
