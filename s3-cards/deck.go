package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Reading From the Hard Drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	// Error handling -> Handled right after
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// rand package, swap, pseudo random generator, randomizing seed
func (d deck) shuffle() {
	// Randomizing the seed
	// Create new Source -> Using current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// Swap positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
