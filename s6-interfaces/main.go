package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// Without interfaces
	printGreetingEB(eb)
	printGreetingSB(sb)

	// With interfaces
	printGreeting(eb)
	printGreeting(sb)
}

// Issue with different types input -> error same function name
// Solution: Change function name, but don't you wish there's an easier solution
func printGreetingEB(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
func printGreetingSB(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

// With Interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
