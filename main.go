package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/redmaner/zipcrypt/src/compress"
	"github.com/redmaner/zipcrypt/src/crypt"
)

func main() {

	args := os.Args
	if len(args) <= 1 {
		os.Exit(1)
	}

	a := compress.NewArchive()
	err := a.Add(args[1])
	if err != nil {
		log.Fatal(err)
	}

	data, err := a.Compress()
	if err != nil {
		log.Fatal(err)
	}

	c := crypt.NewCoffin(crypt.CryptCHACHA20Poly1305)
	c.Opts.WithNonce = true
	cryptData, err := c.Encrypt(data, []byte("This is a simple test"))
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("./test.zcrypt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, bytes.NewBuffer(cryptData))
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.Open("./test.zcrypt")
	if err != nil {
		log.Fatal(err)
	}

	data, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	plaintext, err := c.Decrypt(data, []byte("This is a simple test"))
	if err != nil {
		log.Fatal(err)
	}

	zip, err := os.Create("./test.zip")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(zip, bytes.NewBuffer(plaintext))
	if err != nil {
		log.Fatal(err)
	}
	zip.Close()
}
