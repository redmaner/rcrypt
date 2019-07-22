package main

import "log"

// errorLog is een generieke functie om een fout te loggen
func errorLog(err error, msg string) {
	if err != nil {
		log.Printf("%s: %v", msg, err)
	}
}

// errorPanic is een generieke functie om een fout te loggen en af te sluiten (panic)
func errorPanic(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
