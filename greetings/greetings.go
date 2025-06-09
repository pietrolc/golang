package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func RandomGreeting() (string, error) {
	// List of greetings
	greetings := []string{
		"Hello, World!",
		"Welcome to Go programming!",
		"Hi there, Gopher!",
		"Greetings from the Go team!",
	}

	// Generate a random index
	index := rand.Intn(len(greetings))

	// Return a random greeting
	return greetings[index], nil
}
