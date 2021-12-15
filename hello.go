package main

import (
	"fmt"

	greetings "github.com/runcess/module3_greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
