package main

import "fmt"

// Variable Declarations
func main() {
	// Initialise
	// var card string = "Ace of Spades"
	card := "Ace of Spades" // Go infer that the value is of String type automatically
	// Reassign
	card = "Five of Diamonds" // omit ':'

	fmt.Println(card)
}
