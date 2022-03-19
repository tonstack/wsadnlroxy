package proxy

import (
	"net"
	"net/http"
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

	for {
		select {
		case wsMsg := <-wsChan:
			if wsMsg == nil {
				return // connection closed
			}
			logrus.Debug("wsMsg case: " + string(wsMsg))
			tcpconn.Write(wsMsg)

		case tcpMsg := <-tcpChan:
			if tcpMsg == nil {
				return // connection closed
			}
			logrus.Debug("tcpMsg case: " + string(tcpMsg))
			if err := wsconn.WriteMessage(
				websocket.BinaryMessage, tcpMsg,
			); err != nil {
				break
			}
		}
	}
}
