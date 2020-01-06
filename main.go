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
	fmt.Println("pass")
	query := req.URL.Query()
	bytevalue, _ := json.Marshal(query)
	fmt.Printf("%s\n", bytevalue)

	bytebody := bytes.NewBuffer(bytevalue)
	res, err := http.Post("http://localhost:8080/", "application/json", bytebody)
	if err != nil {
		fmt.Println("Request error :", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("[status] %d\n", res.StatusCode)

	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	sb := string(body)
	fmt.Printf("[body] %v\n", sb)

}
