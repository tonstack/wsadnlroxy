package main

import (
	"wstcproxy/config"
	"wstcproxy/proxy"
)

func main() {
	config.Configure()
	proxy.RunServer()
}
