package config

import (
	"flag"
	"fmt"
	"os"
	"time"
	"wstcproxy/helper"

	"github.com/sirupsen/logrus"
)

type MainConfig struct {
	IP   string
	Port string

	ConnDeadline time.Duration

	WSReadLimit        int64
	WSReadBufferSize   int
	WSWriteBufferSize  int
	WSHandshakeTimeout time.Duration

	TCPBufferSize int
}

var CFG MainConfig

func Configure() {
	CFG = MainConfig{
		ConnDeadline: 30 * time.Second,

		WSReadLimit:        16384,
		WSReadBufferSize:   8192,
		WSWriteBufferSize:  8192,
		WSHandshakeTimeout: 1 * time.Second,

		TCPBufferSize: 8192,
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	var hostport string

	flag.Usage = func() {
		fmt.Print(
			"Usage of wstcproxy:\n\n",

			"--host\tstring\t(required)\n",
			"\tThe host:port on which the server \n",
			"\twill be up. Example: 127.0.0.1:3000\n\n",

			"--debug\tbool\t(optional)\n",
			"\tActivates debug mode. Don't \n",
			"\tuse this flag in production.\n",
		)
		os.Exit(1)
	}

	isDebug := flag.Bool("debug", false, "")
	flag.StringVar(&hostport, "host", "", "")

	flag.Parse()

	if *isDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if hostport == "" {
		flag.Usage()
	}

	var err error

	CFG.IP, CFG.Port, err = helper.SepIPPort(hostport)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
