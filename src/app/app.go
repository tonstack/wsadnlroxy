package app

import (
	"errors"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type mainConfig struct {
	HOST string
	PORT string

	ConnDeadline time.Duration

	WSReadLimit        int64
	WSReadBufferSize   int
	WSWriteBufferSize  int
	WSHandshakeTimeout time.Duration

	TCPBufferSize int
}

var CFG mainConfig

func Configure() error {
	logrus.Info("starting the \"Configure\" function")

	CFG = mainConfig{
		ConnDeadline: 60 * time.Second,

		WSReadLimit:       16384,
		WSReadBufferSize:  16384,
		WSWriteBufferSize: 16384,
		TCPBufferSize:     16384,

		WSHandshakeTimeout: 1 * time.Second,
	}

	CFG.HOST = os.Getenv("APP_HOST")
	CFG.PORT = os.Getenv("APP_PORT")

	if CFG.HOST == "" || CFG.PORT == "" {
		return errors.New("\"HOST\" and \"PORT\" environment vars cannot be empty")
	}

	return nil
}
