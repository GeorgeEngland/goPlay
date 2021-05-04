package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("product", "ETH-USD", "product_id")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "ws-feed.pro.coinbase.com"}
	log.Printf("connecting to %s", u.String())
	c, res, err := websocket.DefaultDialer.Dial(u.String(), nil)
	fmt.Println(res)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	byteValue, _ := ioutil.ReadFile("sub.json")
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result)
	err = c.WriteJSON(result)
	if err != nil {
		log.Fatal("JSON:", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		var m sync.Mutex
		for {
			_, message, err := c.ReadMessage()
			start := time.Now()

			if err != nil {
				log.Println("read:", err)
				return
			}

			go func() {
				var result map[string]interface{}

				err = json.Unmarshal(message, &result)
				msg2 := []byte(" " + strconv.Itoa(int(time.Now().Sub(start).Microseconds())) + " us")
				m.Lock()
				err2 := ioutil.WriteFile("file.txt", msg2, 0644)
				m.Unlock()
				if err2 != nil {
					fmt.Println(err2)
				}
			}()

			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println(elapsed, result)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
		}
	}
}
