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

	fmt.Println("Finished WaitGroup concing")
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}

	}()
	go func() {
		for {
			c2 <- "Every 200ms"
			time.Sleep(time.Millisecond * 200)
		}

	}()

	for {
		fmt.Print("PRINTING: ")
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

func count(thing string, delay time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * delay)
	}
}
