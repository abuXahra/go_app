package main

import (
	"fmt"
)

func main() {

	fmt.Println("=================for loop method 1==============")

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("=================for loop method 2==============")
	j := 1
	for ; j <= 5; j++ {
		fmt.Println(j)
	}
	fmt.Println(j)

	fmt.Println("=================for loop  method 3==============")

	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}

	fmt.Println("=================Break in for loop  method 4==============")

	l := 0

	for {
		fmt.Println(l)
		l++

		if l == 5 {
			break
		}
	}

	fmt.Println("=================Continue in for loop  method 4==============")

	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			continue
		}
		fmt.Println(a)
	}

	fmt.Println("=================nexted forloop in for loop  method 4==============")
	//Loop:
	for n := 1; n <= 3; n++ {
		for m := 1; m <= 3; m++ {
			fmt.Println(n * m)
			if n*m >= 3 {
				break //Loop
			}
		}
	}

}
