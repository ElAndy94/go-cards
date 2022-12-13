package main

import "fmt"

type bot interface {
	getGreeting() string
	// getGreeting(string, int) (string, error)
	// getBotVersion() float64
	// respondToUser(string) string
}
type englishBot struct{}
type spanishBot struct{}

func interfaces() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
