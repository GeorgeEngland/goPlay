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

type l2 struct {
	Changes   [][]string `json:"changes"`
	ProductID string     `json:"product_id"`
	Time      string     `json:"time"`
	Type      string     `json:"type"`
}

type bookItem struct {
	price int64
	quant float32
}

var addr = flag.String("product", "ETH-USD", "product_id")

func main() {
	flag.Parse()
	log.SetFlags(0)
	var askList List
	var bidList List

	//askList.Init()
	//bidList.Init()
	var oBook = Orderbook{&askList, &bidList}

	oBook.askList.append(&Element{Value: bookItem{price: 32, quant: .123}})
	oBook.bidList.append(&Element{Value: bookItem{price: 32, quant: .123}})
	fmt.Println(oBook.askList)

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
			var result l2
			//var result2 map[string]interface{}
			//var s string = fmt.Sprintf("%s", message[9:15])
			err = json.Unmarshal(message, &result)

			go func() {
				if result.Type != "l2update" {
					return
				}
				if err != nil {
					fmt.Println(err)
					return
				}
				msg2 := []byte(" " + strconv.Itoa(int(time.Now().Sub(start).Microseconds())) + " us")
				m.Lock()
				err = ioutil.WriteFile("file.txt", msg2, 0644)
				m.Unlock()
				if err != nil {
					fmt.Println(err)
				}
			}()

			t := time.Now()
			elapsed := t.Sub(start)
			if result.Type == "l2update" {
				res := result.Changes[0]
				fmt.Println(elapsed, res)
			}
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
