package main

import "fmt"

// Slices and For Loops
func main() {
	// Declaring a slice
	cards := []string{newCard(), "Ace of Diamonds"}
	// Adding elements
	cards = append(cards, "Six of Spades") // Does not mutate, returns new slice

	// Iterating over a Slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

// Function
func newCard() string {
	return "Five of Diamonds"
}
