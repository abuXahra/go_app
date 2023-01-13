package main

import (
	"fmt"
)

// func displayError() {
// 	fmt.Sprintln("Therer is no single name call Abdul...\nPlease Input the remaining")
// }

// func phoneNumberError(length string) {
// 	fmt.Sprintln(`your phone number must not be ${length} than 11`)
// }

func main() {
	fmt.Println("===================Control Flow (if, switchcase)==========================")
	fmt.Println("=================if==============")
	number := 50
	guess := 105

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println(" Too low")
		}
		if guess > number {
			fmt.Println(" Too high")
		}
		if guess == number {
			fmt.Println("correct")
		}
	}

}

// variable declarations:
// var decimalNum int
// var firstName string
// var lastName string
// var nameAttachment string
// var address string
// var dob string
// var phoneNumber int

// name := "Abdul"

// fmt.Printf("What is your age:")
// fmt.Scanf("%d", &decimalNum)

// if decimalNum < 0 {
// 	fmt.Println("Invalid input")
// } else if decimalNum > 0 && decimalNum <= 10 {
// 	fmt.Println("First name:")
// 	fmt.Scanf("%s", &firstName)
// 	if firstName == name {
// 		displayError()
// 		fmt.Println(name, " Attachment:")
// 		fmt.Scanf("%s", &nameAttachment)
// 	}

// 	fmt.Println("Last Name:")
// 	fmt.Scanf("%s", &firstName)
// 	if lastName == name {
// 		displayError()
// 		fmt.Println(name, " Attachment:")
// 		fmt.Scanf("%s", &nameAttachment)
// 	}

// 	fmt.Println("Date of Birth:")
// 	fmt.Scanf("%s", &dob)

// 	fmt.Println("Enter Your Phone Number")
// 	fmt.Scanf("%d", &phoneNumber)

// 	fmt.Println("Address:")
// 	fmt.Scanf("%s", address)

// }
