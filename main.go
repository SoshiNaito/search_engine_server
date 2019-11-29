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
	http.HandleFunc("/", jsonhandle)

	http.ListenAndServe(":9080", nil)

}

func jsonhandle(w http.ResponseWriter, req *http.Request) {

	bytevalue, _ := json.Marshal(req.URL.Query())

	fmt.Println(req.URL.Query())

	response, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer(bytevalue))
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	fmt.Printf("[status] %d\n", response.StatusCode)

	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("[body] " + string(body))

}
