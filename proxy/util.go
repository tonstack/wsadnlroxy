package proxy

import (
	"net"
	"wstcproxy/config"

	"github.com/gorilla/websocket"
)

func readTcpBytes(conn net.Conn) ([]byte, error) {
	buf := make([]byte, config.CFG.TCPBufferSize)

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
