package main

import (
	"flag"
	"math/rand"
	"time"

	"sushita_serve/server"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	server.Serve(addr)
}
