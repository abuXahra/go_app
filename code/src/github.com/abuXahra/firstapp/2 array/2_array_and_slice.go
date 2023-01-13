package main

import "fmt"

func arrayAndSlice() {

	fmt.Println("\n\n\t================Array==================\n\n")

	grade1 := 97
	grade2 := 85
	grade3 := 93

	fmt.Printf("\tGrades: %v, %v, %v\n\n", grade1, grade2, grade3)

	grades := [...]int{97, 85, 93}

	fmt.Printf("\tGrades: %v, \n\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[2] = "Ahmed"
	students[1] = "Arnold"
	fmt.Printf("1st Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students)) //length of array

	fmt.Println("=============Array of Array==============")
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// Copying Array
	fmt.Println("=============Copying Array==============")
	a := [...]int{1, 2, 3}
	//b := a // copy all element of a into be
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// Slice of Array
	fmt.Println("\n\n=============Slice Array==============")
	c := []int{1, 2, 3, 4} //slice
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n\n", cap(c))

	c55 := []int{1, 2, 3, 4}
	c55 = append(c55, 45)

	fmt.Println("\n\n=============copy Slice Array==============")
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	e := d[:]   // slice of all elments
	f := d[3:]  // slice from 4th elments to end
	g := d[:6]  // slice first 6 elements
	h := d[3:6] // slice the 4th, 5th, and 6th elments
	d[5] = 45
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println("\n\n=============Using string Arrays make method==============")

	a1 := []string{}
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n\n", cap(a1))
	a1 = append(a1, "1")
	fmt.Println(a1)
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n\n", cap(a1))
	a1 = append(a1, "a", "b", "c", "d", "e")
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n\n", cap(a1))

	c1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	b1 := append(c1[:2], c1[3:]...)
	b2 := c1[3:]
	b3 := c1[:len(c1)-1]
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(c1)

	fmt.Println("\n\n============================Variables=====================")

	// a2 := make([]int, 3, 100)
	// fmt.Printf("Length: %v\n", len(a2))
	// fmt.Printf("Capacity: %v\n\n", cap(a2))

}
