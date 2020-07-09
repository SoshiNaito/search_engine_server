package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	//CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//受け取ったリクエストからクエリを抽出
	getQuery := req.URL.Query()
	queryByteValue, err := json.Marshal(getQuery)
	if err != nil {
		log.Println("error: GET query encode")
		log.Fatal(err)
	}

	//Postリクエスト
	query := bytes.NewBuffer(queryByteValue)
	endPoint := "/search"
	url := fmt.Sprintf("http://localhost:8080%s", endPoint)
	res, err := http.Post(url, "application/json", query)
	if err != nil {
		log.Println("error: POST request")
		log.Fatal(err)
	}
	defer res.Body.Close()

	//レスポンスの読み取り
	resValue, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error: response encode")
		log.Fatal(err)
	}
	w.Write(resValue)
}
