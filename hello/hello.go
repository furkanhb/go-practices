package main

import (
	"fmt"
	"log"

	"example.com/greatings"
)

func main() {
	log.SetPrefix("Greatings: ")
	log.SetFlags(5)

	names := []string{"Furkan", "Abuzer", "Kadayif"}

	messages, err := greatings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
