package main

// Slices and For Loops
func main() {
	// Declaring a slice
	cards := deck{newCard(), "Ace of Diamonds"}
	// Adding elements
	cards = append(cards, "Six of Spades") // Does not mutate, returns new slice

	cards.print()
}

// Function
func newCard() string {
	return "Five of Diamonds"
}
