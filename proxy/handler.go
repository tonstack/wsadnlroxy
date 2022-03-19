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

func readTcpBytes(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)

	size, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:size], nil
}

func chanFromWSConn(wsconn *websocket.Conn) chan []byte {
	wsc := make(chan []byte)

	go func() {
		for {
			_, wsmsg, err := wsconn.ReadMessage()
			if err != nil {
				wsc <- nil
				break
			}
			wsc <- wsmsg
		}
	}()

	return wsc
}

func chanFromTCPConn(tcpconn net.Conn) chan []byte {
	tcpc := make(chan []byte)

	go func() {
		for {
			tcpmsg, err := readTcpBytes(tcpconn)
			if err != nil {
				tcpc <- nil
				break
			}
			tcpc <- tcpmsg
		}
	}()

	return tcpc
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	defer wsconn.Close()

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
