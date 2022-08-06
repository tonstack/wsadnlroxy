package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/tonstack/wsadnlroxy/src/app"
	"github.com/tonstack/wsadnlroxy/src/log"
	"github.com/tonstack/wsadnlroxy/src/proxy"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:      func(r *http.Request) bool { return true },
	ReadBufferSize:   app.CFG.WSReadBufferSize,
	WriteBufferSize:  app.CFG.WSWriteBufferSize,
	HandshakeTimeout: app.CFG.WSHandshakeTimeout,
}

func main() {
	if err := log.Configure(); err != nil {
		panic(err)
	}

	logrus.Info("starting the \"wstcproxy\"")
	if err := app.Configure(); err != nil {
		logrus.Fatalf("can't configure app: %s", err)
	}

	http.HandleFunc("/", proxy.MainHandler)

	hostport := fmt.Sprintf("%s:%s", app.CFG.HOST, app.CFG.PORT)

	logrus.Infof("starting server on: %s", hostport)
	logrus.Fatal(http.ListenAndServe(hostport, nil))
}
