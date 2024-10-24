package main

import (
	"fmt"
	"primer2/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Nigga")
	fmt.Println(message)
}
