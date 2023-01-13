package main

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message
}

func main() {
	fmt.Println(Hello("Abdulmumin"))

	var x int = 100
	var y int = 200
	var b *int = &x
	var c *int = &y

	// z = &x

	// y = *z

	fmt.Println(b, " ", c)
}
