package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/mox692/sushita_serve/client"
	"github.com/mox692/sushita_serve/db"
	"github.com/mox692/sushita_serve/handler"
	"github.com/mox692/sushita_serve/server"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	db.Init()
}

func main() {
	// 第二引数の関数はHandlerのインターフェースの型を満たすものを配置
	client.Get()
	http.HandleFunc("/ranking", handler.GetRanking)
	// ポート8080番でサーバーを起動する
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("listing to serve...")
	}
	server.Serve(addr)

}
