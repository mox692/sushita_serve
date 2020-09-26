package main

import (
	"fmt"
	"net/http"

	"../client"
	"../handler"
)

func main() {
	// 第二引数の関数はHandlerのインターフェースの型を満たすものを配置
	client.Get()
	// http.HandleFunc("/ranking", handler.RankingGet)
	http.HandleFunc("/ranking", handler.GetRanking)
	// ポート8080番でサーバーを起動する
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("listing to serve...")
	}

}
