package main

import (
	"fmt"

	"github.com/pietrolc/golang/greetings" // Replace with your actual module path
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Pietro")
	fmt.Println(message)
}
