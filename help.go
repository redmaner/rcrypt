package main

import "fmt"

const usage = `
Usage: rcrypt [command] [options] file(s)...
Version: %s

Commands:
    help              This help
    seal              Encrypt file(s)
    open              Decrypt file(s)

`

const usageSeal = `
usage: rcrypt seal [options] file(s)...

Options:
    --help, -h            This help
    --nonce, -n           Use nonce for encryption (applies to AES and chacha20 only)
    --out, -o             The path to encrypt the data
    --passin, -p          The password used to encrypt the data

`

const usageOpen = `
usage: rcrypt open [options] file(s)...

Options:
    --help, -h            This help
    --nonce, -n           Path to nonce, if nonce was used for encryption (applies to AES and chacha20 only)
    --out, -o             The path to decrypt the data
    --passin, -p          The password used to decrypt the data

`

func showHelp() {
	fmt.Printf(usage, "0.1.0")
}

func showHelpSeal() {
	fmt.Printf(usageSeal)
}

func showHelpOpen() {
	fmt.Printf(usageOpen)
}
