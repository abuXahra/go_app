package main

import (
	"fmt"
)




func main() {


	// var name string
    // var alphabet_count int
    // var float_value float32
    // var bool_value bool
 

	var name string
	var alphabet_count int

	fmt.Scan(&name)
	fmt.Scan(&alphabet_count)

	fmt.Printf("The word %s containing %d number of alphabets", name, alphabet_count)

	// fmt.Printf("Number of Apple?")
	// fmt.Scan("%d", &appleNum)
	// fmt.Printf("%d", appleNum)
}
