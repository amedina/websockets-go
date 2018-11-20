package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize:1024,
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Echo handler...")
	conn, _ := upgrader.Upgrade(w, r, nil)
	
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to the browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func main() {
	fmt.Println("Websockets...")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/echo", echoHandler)

	http.ListenAndServe(":8080", nil)
}