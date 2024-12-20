package main

import (
	"fmt"
	"log"
	"primer2/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("Holy shit: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Аллах")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
