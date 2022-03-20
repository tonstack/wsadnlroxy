package proxy

import (
	"net"
	"net/http"
	"time"
	"wstcproxy/config"
	"wstcproxy/helper"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer wsconn.Close()
	wsconn.SetReadLimit(config.CFG.WSReadLimit)

	hostport := r.URL.Query().Get("dest_host")
	if _, _, err = helper.SepIPPort(hostport); err != nil {
		wsconn.Close()
		return
	}

	tcpconn, err := net.Dial("tcp", hostport)
	if err != nil {
		wsconn.Close()
		return
	}

	tcpChan := chanFromTCPConn(tcpconn)
	wsChan := chanFromWSConn(wsconn)

	deadline := time.NewTicker(3 * time.Second)
	var lastmsg time.Time = time.Now()

	for {
		select {
		case wsMsg := <-wsChan:
			lastmsg = time.Now()

			if wsMsg == nil {
				logrus.Debug("ws conn closed")
				return
			}
			logrus.Debug("wsMsg case: " + string(wsMsg))
			tcpconn.Write(wsMsg)

		case tcpMsg := <-tcpChan:
			lastmsg = time.Now()

			if tcpMsg == nil {
				logrus.Debug("tcp conn closed")
				return
			}
			logrus.Debug("tcpMsg case: " + string(tcpMsg))
			if err := wsconn.WriteMessage(websocket.BinaryMessage, tcpMsg); err != nil {
				return
			}

		case <-deadline.C:
			logrus.Debug("check deadline")
			if time.Now().Sub(lastmsg) > config.CFG.ConnDeadline {
				logrus.Debug("end deadline")
				return
			}
		}
	}
}
