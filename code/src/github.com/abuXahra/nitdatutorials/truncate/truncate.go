package main

import (
	"fmt"
)

func displayError() {
	fmt.Sprintln("Therer is no single name call Abdul...\nPlease Input the remaining")
}

func phoneNumberError(length string) {
	fmt.Sprintln(`your phone number must not be ${length} than 11`)
}

func main() {
	var decimalNum float32
	var firstName string
	var lastName string
	var nameAttachment string
	var address string
	var dob string
	var phoneNumber int
	name := "Abdul"

	fmt.Printf("What is your age:")

	if decimalNum < 0 {
		fmt.Println("Invalid input")
	} else if decimalNum > 0 && decimalNum <= 10 {
		fmt.Println("First name:\n")
		fmt.Scanf("%s", &firstName)
		if firstName == name {
			displayError()
			fmt.Scanf("%s", &nameAttachment)
		}

		fmt.Println("Last Name:\n")
		fmt.Scanf("%s", &firstName)
		if lastName == name {
			displayError()
			fmt.Scanf("%s", &nameAttachment)
		}

		fmt.Println("Date of Birth:\n")
		fmt.Scanf("%s", &dob)

		fmt.Println("Enter Your Phone Number\n")
		fmt.Scanf("%d", &phoneNumber)

	}

	fmt.Scanf("%f", &decimalNum)

	var result = int(decimalNum)

	fmt.Printf("%d", result)
	fmt.Printf("%s", address)
}
