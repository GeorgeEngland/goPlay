package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "ws-feed.pro.coinbase.com", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	byteValue, _ := ioutil.ReadFile("sub.json")

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	b := bytes.NewBuffer(byteValue)
	//json.NewEncoder(b).Encode(result)
	fmt.Println(result, b)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr}
	log.Printf("connecting to %s", u.String())
	c, res, err := websocket.DefaultDialer.Dial(u.String(), nil)
	fmt.Println(res)
	if err != nil {
		log.Fatal("dial:", err)
	}

	err = c.WriteJSON(result)
	if err != nil {
		log.Fatal("JSON:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {

			_, message, err := c.ReadMessage()
			start := time.Now()

			if err != nil {
				log.Println("read:", err)
				return
			}
			var result map[string]interface{}

			err = json.Unmarshal(message, &result)
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
