package proxy

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	ID   string
	Data string
}

func readTcpBytes(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)

	size, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:size], nil
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer wsconn.Close()

	// TODO: parse ip with port
	dsthost := r.URL.Query().Get("dstip")
	// if net.ParseIP(dstip) == nil {
	// 	wsconn.Close() // TODO: send msg with err
	// 	return
	// }127.0.0.1:8010

	tcpconn, err := net.Dial("tcp", dsthost)
	if err != nil {
		wsconn.Close() // TODO: send msg with err
		return
	}

	for {
		_, wsmsg, err := wsconn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		tcpconn.Write(wsmsg)

		data, err := readTcpBytes(tcpconn)
		log.Println("data:", string(data))

		err = wsconn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
