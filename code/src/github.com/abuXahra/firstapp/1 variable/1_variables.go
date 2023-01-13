package main

import "fmt"

//block variable declarations
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "sarah Jane smith"
	doctorNumber        = 3
	season       int    = 11
)

var (
	counter int = 0
)

const a1 int = 100

const (
	r  = iota
	ra = iota
	rb = iota
)

func Variables() {

	name := "Isah Abdulmumin"

	nameText := "String Data Type"
	fmt.Println(nameText)
	fmt.Printf("%v, %T", name, name)

	fmt.Println("\n\t => Bulk Variable Declaration")
	fmt.Println(actorName)
	fmt.Println(companion)
	fmt.Println(doctorNumber)
	fmt.Println(season)

	fmt.Println("\n\t => boolean data types")
	n := 1 == 1
	m := 1 == 2

	var c bool
	fmt.Println(c)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n\n", m, m)

	fmt.Println("\t => integer data types")

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Println("\t => complex data types")

	d := 1 + 2i
	e := 2 + 4.3i
	fmt.Println(d + e)
	fmt.Println(d - e)
	fmt.Println(d * e)
	fmt.Println(d / e)

	fmt.Println("\n\t => String data types\n")

	s := "this is a string\n"
	p := []byte(s)
	fmt.Printf("%v, %T", s, s)
	fmt.Printf("%v, %T", p, p)
	fmt.Printf("\n")

	fmt.Println("\n\t => Rune data types\n")
	var q rune = 23

	fmt.Println(q)
	fmt.Printf("%v, %T\n", q, q)

	fmt.Println("\n\t => Constants\n")
	const a1 int = 14
	const b1 string = "foo"
	const c1 float32 = 3.14
	const d1 bool = true

	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", c1)
	fmt.Printf("%v\n", d1)

	fmt.Printf("%v\n", r)
	fmt.Printf("%v\n", ra)
	fmt.Printf("%v\n", rb)

	fmt.Println("\n\t => Array\n")

}
