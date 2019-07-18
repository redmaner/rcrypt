package main

import (
	"log"
	"os"

	"github.com/redmaner/zipcrypt/src/compress"
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
	a.CompressToFile("test.zip")
	if err != nil {
		log.Fatal(err)
	}
}
