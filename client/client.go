package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Get() {
	// url := "http://google.co.jp"
	// resp, err := http.Get(url)

	client := new(http.Client)
	req, err := http.NewRequest("GET", "http://example.com", nil)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray)) // htmlをstringで取得
}

func Post() {

}
