package proxy

import (
	"log"
	"net"
	"net/http"
	"wstcproxy/helper"

	"github.com/gorilla/websocket"
)

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

	for {
		_, wsmsg, err := wsconn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		tcpconn.Write(wsmsg)

		data, err := readTcpBytes(tcpconn)

		err = wsconn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			break
		}
	}
}
