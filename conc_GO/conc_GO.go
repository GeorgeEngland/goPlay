package main

import (
	"fmt"
	"log"

	"play.com/greetings_GO"
)

func main() {
	fmt.Println("testing conc")
	msg, err := greetings_GO.Hello("George")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
