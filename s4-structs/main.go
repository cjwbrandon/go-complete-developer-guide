package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // Embedding structs
}

func main() {
	// Assigning structs -> assigning by order of fields
	// alex := person{"Alex", "Anderson"}
	// // Another way -> more explicit
	// peter := person{firstName: "Peter", lastName: "Anderson"}
	// // 3rd way -> assigned zero value
	// var bob person           // both empty strings
	// fmt.Printf("%+v\n", bob) // %+v -> Print out field names & values

	// // Updating Struct Values
	// bob.firstName = "Bob"
	// fmt.Println(alex)
	// alex.firstName = "Alexx"

	// fmt.Println(alex, peter)

	// Embedded Structs
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim // Get reference to memory address
	// jimPointer.updateName("Jimmy")
	// Shortcut
	jim.updateName("Jimmy")
	jim.print()

	// Gotchas with Pointers
	// Slices get updated
	random()
}

// *type -> pointerToType
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer -> get value
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func random() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
