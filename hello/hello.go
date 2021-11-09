package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// message, err := greetings.Hello("OBI Wycliffe")
	message, err := greetings.Hello(randomName())
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of names. The returned
// message is selected at random.
func randomName() string {
	// A slice of names.
	names := []string{
		"OBI Wycliffe",
		"Lucy Amondi",
		"Ogutu Kennedy",
		"Nick Wick",
	}

	// Return a randomly selected name by specifying
	// a random index for the slice of formats.
	return names[rand.Intn(len(names))]
}
