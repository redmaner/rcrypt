package main

import "fmt"

func getPassword() string {
	var passwordOne string
	var passwordTwo string

	fmt.Printf("Enter password: ")
	if _, err := fmt.Scan(&passwordOne); err != nil {
		fmt.Printf("\nSomehting went wrong\n\n")
		return getPassword()
	}

	fmt.Printf("Re-enter password: ")
	if _, err := fmt.Scan(&passwordTwo); err != nil {
		fmt.Printf("\nSomehting went wrong\n\n")
		return getPassword()
	}

	if passwordOne != passwordTwo {
		fmt.Printf("\nPasswords didn't match\n")
		return getPassword()
	}

	return passwordOne
}
