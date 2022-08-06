package proxy

import (
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/tonstack/wsadnlroxy/src/app"
	"github.com/tonstack/wsadnlroxy/src/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:      func(r *http.Request) bool { return true },
	ReadBufferSize:   app.CFG.WSReadBufferSize,
	WriteBufferSize:  app.CFG.WSWriteBufferSize,
	HandshakeTimeout: app.CFG.WSHandshakeTimeout,
}

type liteServerConfig struct {
	hostport string
	pubkey   string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("new client connected")

	respHeader := make(http.Header)
	respHeader.Add("Wsadnlroxy-Version", app.VERSION)

	wsconn, err := upgrader.Upgrade(w, r, respHeader)
	if err != nil {
		return
	}

	defer wsconn.Close()
	wsconn.SetReadLimit(app.CFG.WSReadLimit)

	lsparams, err := readGetArgs(r.URL.Query())
	if err != nil {
		logrus.Debug(err)
		return
	}

	if err := checkLiteServerConnection(lsparams); err != nil {
		logrus.Debug(err)
		return
	}

	tcpconn, err := net.Dial("tcp", lsparams.hostport)
	if err != nil {
		logrus.Debug(err)
		return
	}

	tcpChan := utils.ChanFromTCPConn(tcpconn, app.CFG.TCPBufferSize)
	wsChan := utils.ChanFromWSConn(wsconn)

	checkDeadlineTicker := time.NewTicker(1 * time.Second)
	var lastmsg time.Time = time.Now()

	for {
		select {
		case wsMsg := <-wsChan:
			lastmsg = time.Now()

			if wsMsg == nil {
				logrus.Debug("ws conn empty - continue")
				continue
			}

			logrus.Debugf("new ws msg recv bytes: %d", len(wsMsg))
			tcpconn.Write(wsMsg)

		case tcpMsg := <-tcpChan:
			lastmsg = time.Now()

			if tcpMsg == nil {
				logrus.Debug("tcp conn empty - continue")
				continue
			}

			logrus.Debugf("new tcp msg recv bytes: %d", len(tcpMsg))
			if err := wsconn.WriteMessage(websocket.BinaryMessage, tcpMsg); err != nil {
				return
			}

		case <-checkDeadlineTicker.C:
			since, dead := time.Since(lastmsg), app.CFG.ConnDeadline

			logrus.Debugf(
				"check deadline (since: %ds; dead: %ds)",
				since/time.Second, dead/time.Second,
			)

			if since >= dead {
				logrus.Debug("end deadline")
				return
			}
		}
	}
}
