package config

import (
	"flag"
	"fmt"
	"os"
	"wstcproxy/helper"

	log "github.com/sirupsen/logrus"
)

type MainConfig struct {
	IP   string
	Port string
}

var CFG MainConfig

func Configure() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

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
		log.SetLevel(log.DebugLevel)
	}
	if hostport == "" {
		flag.Usage()
	}

	var err error

	CFG.IP, CFG.Port, err = helper.SepIPPort(hostport)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
