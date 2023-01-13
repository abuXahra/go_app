package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number    int
	actorName string
	companion []string
	episode   []string
}

// /similar to inheritance because there is not inheritance in go programing
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {

	//Struct
	fmt.Println("================Struct function in go")
	fmt.Println("================1. First Method of Declaration")
	aDoctor := Doctor{
		number:    3,
		actorName: "idris fatimah",
		companion: []string{"ggh", "jkhhk"},
		episode:   []string{"one", "two", "three"},
	}
	fmt.Println(">>>> ", aDoctor, " <<<<<< ")

	fmt.Println("================2. Second Method of Declaration===================================")
	bDoctor := struct{ name string }{name: "pat Attah"}

	fmt.Println(">>>> ", bDoctor, " <<<<<< ")

	fmt.Println("================3. Second Method of Declaration==============================")
	type University struct {
		uniId      int
		uniName    string
		uniDesc    string
		uniAddress string
	}

	myUniversity := University{
		uniId:      2,
		uniName:    "Federal University of Technology Minna",
		uniDesc:    "University that deals with technological courses",
		uniAddress: "Federal University of Technology, Gidan Kwano, Minna - Niger State",
	}

	fmt.Println(myUniversity.uniAddress)

	fmt.Println("================3. Second Method of Declaration==============================")
	type Names struct {
		name   string
		number int
		detail []string
	}

	sNames := Names{}
	sNames.name = "isah abdulmumin"
	sNames.number = 45
	sNames.detail = []string{
		"Plot 107 jeguin estate mabushi abuja\n",
		"Plot 107 jeguin estate mabushi abuja\n",
		"Plot 107 jeguin estate mabushi abuja\n",
	}
	fmt.Println(aDoctor, "\n", sNames)

	//similar to iharitance printout
	bird := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48.90,
		CanFly:   false,
	}

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	fmt.Println(bird, "\n", t)

}
