package main

import (
	"fmt"

	"example.com/greetings"

	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Zevik", "Vy heo", "Ngoc thuc"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Printf("%s: %s\n", name, message)
	}
}
