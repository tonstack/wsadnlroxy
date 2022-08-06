package utils

import (
	"net"

	"github.com/gorilla/websocket"
)

func readTcpBytes(conn net.Conn, bufferSize int) ([]byte, error) {
	buf := make([]byte, bufferSize)

	size, err := conn.Read(buf)

	if err != nil {
		return nil, err
	}

	return buf[:size], nil
}

func ChanFromWSConn(wsconn *websocket.Conn) chan []byte {
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

func ChanFromTCPConn(tcpconn net.Conn, bufferSize int) chan []byte {
	tcpc := make(chan []byte)

	go func() {
		for {
			tcpmsg, err := readTcpBytes(tcpconn, bufferSize)
			if err != nil {
				tcpc <- nil
				break
			}
			tcpc <- tcpmsg
		}
	}()

	return tcpc
}
