package main

import (
	"flag"
	"math/rand"
	"time"

	"sushita_serve/server"

	"google.golang.org/appengine"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	if !(appengine.IsAppEngine()) {
		flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
		flag.Parse()
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	server.Serve(addr)
}
