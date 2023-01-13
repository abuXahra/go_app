package main

import (
	"fmt"
)

func main() {

	number := 50

	fmt.Println("=================switch case 1==============")
	switch number {
	case 1, 5, 10:
		fmt.Println("one, five ten")

	case 2, 4, 6:
		fmt.Println("two, four or six")

	default:
		fmt.Println("another number")

		fmt.Println("=================switch case 1==============")

		var i interface{} = 1
		switch i.(type) {
		case int:
			fmt.Println("i is an int")
		case float64:
			fmt.Println("i is an float64")
		case string:
			fmt.Println("i is an string")
		default:
			fmt.Println("i is another type")

		}

	}
}
