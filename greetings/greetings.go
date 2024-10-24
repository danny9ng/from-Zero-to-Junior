package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greetings for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name, nigga")
	}

	// If a name was received, return a value that ambeds the name
	// in a greetings message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
