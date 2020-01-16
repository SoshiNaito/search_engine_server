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

	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := req.URL.Query()
	// cquery := pack.Rename("query", query)
	// fmt.Printf("%s\n", cquery)
	// bytevalue, _ := json.Marshal(cquery)
	bytevalue, _ := json.Marshal(query)
	fmt.Printf("%s\n", bytevalue)
	bufbody := bytes.NewBuffer(bytevalue)

	res, err := http.Post("http://localhost:8080/", "application/json", bufbody)
	if err != nil {
		fmt.Println("Request error :", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))

	w.Write(body)
}
