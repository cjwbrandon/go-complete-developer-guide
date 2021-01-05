package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type card struct {
	suit  string
	value string
}

func main() {
	// Assigning structs -> assigning by order of fields
	alex := person{"Alex", "Anderson"}
	// Another way -> more explicit
	peter := person{firstName: "Peter", lastName: "Anderson"}
	// 3rd way -> assigned zero value
	var bob person           // both empty strings
	fmt.Printf("%+v\n", bob) // %+v -> Print out field names & values

	// Updating Struct Values
	bob.firstName = "Bob"
	fmt.Println(alex)
	alex.firstName = "Alexx"

	fmt.Println(alex, peter)
}
