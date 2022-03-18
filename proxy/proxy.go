package proxy

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"wstcproxy/config"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:      func(r *http.Request) bool { return true },
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 1 * time.Second,
}

func RunServer() {
	http.HandleFunc("/", mainHandler)

	hostportString := fmt.Sprintf("%s:%s", config.CFG.IP, config.CFG.Port)
	log.Println("Starting server on:", hostportString)

	log.Fatal(http.ListenAndServe(hostportString, nil))
}
