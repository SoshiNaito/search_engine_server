package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := req.URL.Query()
	bytevalue, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
	}
	bufbody := bytes.NewBuffer(bytevalue)
	res, err := http.Post("http://localhost:8080/", "application/json", bufbody)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	w.Write(body)
}
