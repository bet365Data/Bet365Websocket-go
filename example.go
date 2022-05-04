package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

var (
	addr   = flag.String("addr", "www.bet365data.com", "http service address")
	token  = flag.String("token", "", "your token")
	path   = flag.String("path", "/ws/inplayFootBall/en", "your path") //en=English|cn =Chinese
	dialer *websocket.Dialer
)

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: *path, RawQuery: fmt.Sprintf("token=%s", *token)}
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		if gjson.GetBytes(message, "code").Int() == 0 {
			fmt.Println(gjson.GetBytes(message, "channel")) //Inplay
			fmt.Println(gjson.GetBytes(message, "typs"))    //Event|List|Detail
			fmt.Println(gjson.GetBytes(message, "data"))    //map
		}
	}
}
