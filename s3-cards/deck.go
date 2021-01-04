package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string // Analogy: Extends slice of string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Declare function with receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function with arguments
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Type Conversion + strings package
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saving Data to the Hard Drive
func (d deck) saveToFile(filename string) error {
	// Last argument -> Set permission for the file -> 0666 means anyone can read
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
