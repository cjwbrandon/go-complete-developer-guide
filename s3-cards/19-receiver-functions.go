package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string // Analogy: Extends slice of string

// Declare function with receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func main() {
	cards := deck{"Ace of Diamonds", "Ace of Diamonds"}

	// Executing the function
	cards.print()
}
