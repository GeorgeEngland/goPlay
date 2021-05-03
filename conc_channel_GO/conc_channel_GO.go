package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go count("sheep", 100, c)
	for msg := range c {
		println(msg)
	}

}

func count(thing string, delay time.Duration, c chan string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * delay)
	}
	close(c)
}
