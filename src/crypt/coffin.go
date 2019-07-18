package crypt

type Coffin struct {
	Mode CryptMode
}

func NewCoffin(cm CryptMode) *Coffin {
	return &Coffin{
		Mode: cm,
	}
}
