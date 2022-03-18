package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"wstcproxy/helper"
)

type MainConfig struct {
	Host string
	Port string
}

var CFG MainConfig

func Configure() {
	var hostport string

	flag.Usage = func() {
		fmt.Print(
			"Usage of wstcproxy:\n\n",
			"--host\tstring\n",
			"\tThe host:port on which the server will be up.\n",
			"\tExample: 127.0.0.1:3000\n",
		)
		os.Exit(1)
	}

	hostHelp := "The host:port on which the server will be up. Example: 127.0.0.1:3000"
	flag.StringVar(&hostport, "host", "", hostHelp)
	flag.Parse()

	if hostport == "" {
		flag.Usage()
	}

	if err := helper.SepIPPort(hostport, &CFG.Host, &CFG.Port); err != nil {
		log.Fatalln(err.Error())
	}
}
