package main

import (
	"fmt"
	"log"

	"github.com/pietrolc/golang/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("tester")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// Request a random greeting message.
	message2, err2 := greetings.RandomGreeting()
	// If an error was returned, print it to the console and
	// exit the program.
	if err2 != nil {
		log.Fatal(err2)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message2)
}
