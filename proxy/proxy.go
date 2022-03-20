package proxy

import (
	"fmt"
	"net/http"
	"wstcproxy/config"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:      func(r *http.Request) bool { return true },
	ReadBufferSize:   config.CFG.WSReadBufferSize,
	WriteBufferSize:  config.CFG.WSWriteBufferSize,
	HandshakeTimeout: config.CFG.WSHandshakeTimeout,
}

func RunServer() {
	http.HandleFunc("/", mainHandler)

	hostportString := fmt.Sprintf("%s:%s", config.CFG.IP, config.CFG.Port)
	logrus.Info("Starting server on:", hostportString)

	logrus.Fatal(http.ListenAndServe(hostportString, nil))
}
