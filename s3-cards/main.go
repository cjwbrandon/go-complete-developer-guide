package main

// Slices and For Loops
func main() {
	// Declaring a slice
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
