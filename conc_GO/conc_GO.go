package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"play.com/greetings_GO"
)

func main() {
	msg, err := greetings_GO.Hello("George")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	go count("sheep", 100)
	count("fish", 200)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("poop", 100)
		wg.Done()
	}()
	wg.Wait()
}

func count(thing string, delay time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * delay)
	}
}
