package main

import (
	"fmt"

	"play.com/greetings_GO"
)

func main() {
	message := greetings_GO.Hello("George")
	fmt.Println(message)
}
