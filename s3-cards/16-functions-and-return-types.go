package main

import "fmt"

// Functions and Return Types
func main() {
	card := newCard()

	fmt.Println(card)
}

// Function
func newCard() string {
	return "Five of Diamonds"
}
