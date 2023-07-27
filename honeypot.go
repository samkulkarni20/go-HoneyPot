package main

import (
	"github.com/samkulkarni20/go-HoneyPot/config"
	"github.com/samkulkarni20/go-HoneyPot/tcp"
)

func main() {
	cfg := config.Read()
	tcpServer := tcp.NewServer(cfg.TCP.Ports)
	tcpServer.Start()
}
