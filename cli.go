package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/redmaner/rcrypt/src/compress"
	"github.com/redmaner/rcrypt/src/crypt"
)

var (
	// Commands
	cmdSeal = flag.NewFlagSet("seal", flag.ExitOnError)
	cmdOpen = flag.NewFlagSet("open", flag.ExitOnError)

	// General arguments
	argCHACHA   bool
	argAES      bool
	argOut      string
	argPassword string
	argHelp     bool

	// Specific seal options
	argWithNonce bool

	// Specific open options
	argNonce string
)

func init() {

	// Init seal options
	cmdSeal.BoolVar(&argCHACHA, "chacha20", true, "Use the CHACHA20-Poly1305 algorithm")
	cmdSeal.BoolVar(&argAES, "aes256", false, "Use the AES256-GCM algorithm")
	cmdSeal.BoolVar(&argHelp, "help", false, "Show help")                                       // Long argument
	cmdSeal.BoolVar(&argHelp, "h", false, "Show help")                                          // Short argument
	cmdSeal.StringVar(&argOut, "out", "./", "Destination of the output")                        // Long argument
	cmdSeal.StringVar(&argOut, "o", "./", "Destination of the output")                          // Short argument
	cmdSeal.BoolVar(&argWithNonce, "nonce", false, "Use a nonce for encryption and decryption") // Long argument
	cmdSeal.BoolVar(&argWithNonce, "n", false, "Use a nonce for encryption and decryption")     // Short argument
	cmdSeal.StringVar(&argPassword, "passin", "", "Password used for encryption")               // Long argument
	cmdSeal.StringVar(&argPassword, "p", "", "Password used for encryption")                    // short argument

	// Init open options
	cmdOpen.BoolVar(&argCHACHA, "chacha20", true, "Use the CHACHA20-Poly1305 algorithm")
	cmdOpen.BoolVar(&argAES, "aes256", false, "Use the AES256-GCM algorithm")
	cmdOpen.BoolVar(&argHelp, "h", false, "Show help")                            // Short argument
	cmdOpen.BoolVar(&argHelp, "help", false, "Show help")                         // Long argument
	cmdOpen.StringVar(&argOut, "out", "", "Destination of the output")            // Long argument
	cmdOpen.StringVar(&argOut, "o", "", "Destination of the output")              // Short argument
	cmdOpen.StringVar(&argNonce, "nonce", "", "Path to nonce for decryption")     // Long argument
	cmdOpen.StringVar(&argNonce, "n", "", "Path to nonce for decryption")         // Short argument
	cmdOpen.StringVar(&argPassword, "passin", "", "Password used for decryption") // Long argument
	cmdOpen.StringVar(&argPassword, "p", "", "Password used for decryption")      // short argument
}

/*************************
* CLI Interface
**************************/
func main() {

	// Assign arguments
	args := os.Args
	if len(args) < 2 {
		showHelp()
		os.Exit(1)
	}

	// Switch on command
	switch args[1] {

	/*************************
	* SEAL COMMAND
	**************************/
	case "seal":

		// Parse seal flags
		err := cmdSeal.Parse(args[2:])
		errorPanic(err, "Parsing arguments failed")

		// Make variables
		fileArgs := cmdSeal.Args()
		if len(fileArgs) == 0 || argHelp {
			showHelpSeal()
			os.Exit(1)
		}

		t := time.Now()
		defaultFileName := fmt.Sprintf("./%s.rcrypt", t.Format("20060102150405"))

		// Validate arguments
		fileOut := argOut
		password := argPassword

		// Check output argument
		if fileOut == "./" {
			fileOut = fileOut + defaultFileName
		}

		// Check password argument
		if password == "" {
			password = getPassword()
		}

		// Create a new archive
		archive := compress.NewArchive()

		// Add files to archive
		for _, f := range fileArgs {
			err := archive.Add(f)
			errorPanic(err, "Error when compressing files")
		}

		// Compress data to memory
		data, err := archive.Compress()
		errorPanic(err, "Error when compressing files")

		// Make a new coffin
		var cof *crypt.Coffin

		// Encrypt according to selected algorith. The default algorithm is
		// CHACHA20-poly1305
		switch {
		case argAES:
			cof = crypt.NewCoffin(crypt.CryptAES256)
		default:
			cof = crypt.NewCoffin(crypt.CryptCHACHA20)
		}

		// Do specific configurations for AES and CHACHA
		if argAES || argCHACHA {
			cof.Opts.Password = []byte(password)

			// If nonce is enabled, enable nonce
			if argWithNonce {
				cof.Opts.WithNonce = true
			}
		}

		// Encrypt the data
		encryptedData, err := cof.Encrypt(data)
		errorPanic(err, "Error when encrypting data")

		// Create file on disk
		file, err := os.Create(fileOut)
		errorPanic(err, fmt.Sprintf("Error when creating %s", fileOut))

		// Write encrypted data to file
		_, err = io.Copy(file, bytes.NewBuffer(encryptedData))
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		errorPanic(err, fmt.Sprintf("Error closing %s", fileOut))

		// If nonce is enabled, write nonce to disk
		if argWithNonce {

			// get nonce
			nonce := cof.GetNonce()
			nonceOut := fileOut + ".nonce"

			// Create file
			file, err := os.Create(nonceOut)
			errorPanic(err, fmt.Sprintf("Error when creating %s", nonceOut))

			// Write nonce to file
			_, err = io.Copy(file, bytes.NewBuffer(nonce))
			if err != nil {
				log.Fatal(err)
			}
			err = file.Close()
			errorPanic(err, fmt.Sprintf("Error closing %s", nonceOut))
		}

		/*************************
		 * OPEN COMMAND
		 **************************/
	case "open":

		// Parse open flags
		err := cmdOpen.Parse(args[2:])
		errorPanic(err, "Parsing arguments failed")

		// Make variables
		fileArgs := cmdOpen.Args()
		if len(fileArgs) != 1 || argHelp {
			showHelpOpen()
			os.Exit(1)
		}

		// Validate arguments
		password := argPassword
		outDir := argOut

		// check out argument
		if outDir == "" {
			outDir = fileArgs[0] + "-out"
		}

		// Check password argument
		if password == "" {
			password = getPassword()
		}

		// Open encrypted file
		file, err := os.Open(fileArgs[0])
		errorPanic(err, fmt.Sprintf("Error opening %s", fileArgs[0]))

		// Read data
		encryptedData, err := ioutil.ReadAll(file)
		errorPanic(err, fmt.Sprintf("Error reading %s", fileArgs[0]))

		// Decrypt the data acoording to selecte algorithm. The default algorithm
		// is CHACHA20-poly1305
		var cof *crypt.Coffin

		switch {
		case argAES:
			cof = crypt.NewCoffin(crypt.CryptAES256)
		default:
			cof = crypt.NewCoffin(crypt.CryptCHACHA20)
		}

		// Do specific configurations for AES256 and CHACHA20
		if argAES || argCHACHA {
			cof.Opts.Password = []byte(password)

			// Get nonce if it is defined
			if argNonce != "" {
				nonceFile, err := os.Open(argNonce)
				errorPanic(err, fmt.Sprintf("Error opening %s", argNonce))

				// Read data
				nonce, err := ioutil.ReadAll(nonceFile)
				errorPanic(err, fmt.Sprintf("Error reading %s", argNonce))

				// Assign nonce to coffin
				cof.Opts.WithNonce = true
				cof.Opts.Nonce = nonce
			}
		}

		plaintext, err := cof.Decrypt(encryptedData)
		errorPanic(err, "Error decrypting data")

		// Create a new archive
		archive, err := compress.LoadArchive(plaintext)
		errorPanic(err, "Error extracting data")

		err = archive.Decompress(outDir)
		errorPanic(err, "Error extracting data")

	default:
		showHelp()
		os.Exit(0)
	}
}

/*
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

	c := crypt.NewCoffin(crypt.CryptCHACHA20)
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

	arc, err := compress.LoadArchive(plaintext)
	if err != nil {
		log.Fatal(err)
	}

	err = arc.Decompress("./test-zip-out")
	if err != nil {
		log.Fatal(err)
	}
}
*/
