package main

import "fmt"

func Mapfunction() {

	//Map
	fmt.Println("================Map function in go")
	statePopulations := map[string]int{
		"Abuja":  1000,
		"Kaduna": 2000,
		"Lagos":  100000,
		"Kano":   2000000,
	}

	fmt.Println(statePopulations["kano"])

	kogiLocalGovernment := map[string]int{
		"Wuse":    349,
		"Mabushi": 567,
	}
	//Adding to the maps
	kogiLocalGovernment["jos"] = 90000
	//delete from the map
	delete(kogiLocalGovernment, "jos")
	fmt.Println(kogiLocalGovernment)
	fmt.Println(len(statePopulations))
	sp := statePopulations
	sp["Ohio"] = 1020202020
	fmt.Println("============sp==============\n")
	fmt.Println(sp)
	fmt.Println("============Ohio Value==============\n")
	fmt.Println(sp["Ohio"])
	delete(sp, "Ohio")
	fmt.Println("============after ohio is deleted==============")
	fmt.Println(sp)

	jamaaPositions := map[string]string{
		"Sheik Harrun": "Amir",
	}

	jamaaPositions["Suleiman"] = "Treasurer"
	jamaaPositions["Iliya"] = "Electrician"
	jamaaPositions["Lukman"] = "qori"

	fmt.Println("================Jamaaaaaaaaaaaaaaaa==============")
	fmt.Println(jamaaPositions)
	fmt.Println(len(jamaaPositions))
	pop, ok := jamaaPositions["Lukman"]
	fmt.Println("=========pop========")
	fmt.Println(pop, ok)
	//Struct
	fmt.Println("================Struct function in go")

}
