package main

import (
	"fmt"
	"log"
	"time"

	"play.com/greetings_GO"
)

func main() {
	msg, err := greetings_GO.Hello("George")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	go count("sheep", 1000)
	count("fish", 200)
}

func count(thing string, delay time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * delay)
	}
}
