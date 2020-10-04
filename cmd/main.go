package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
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
	// http.HandleFunc("/ranking", handler.GetRanking)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("listing to serve...")
	}
	server.Serve(addr)

}
