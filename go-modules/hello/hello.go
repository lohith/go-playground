package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("hello: ")
	log.SetFlags(0)

	log.Println("Starting")
	msg, err := greetings.Hello("Atharv")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

	log.Println()

	names := []string{"Shilpa", "Akshobya", "Lohith"}

	log.Println(greetings.Hellos(names))

}
