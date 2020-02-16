package main

import (
	"bytes"
	"io/ioutil"

	//"bytes"
	"encoding/json"
	"fmt"
	"time"

	//"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)

}

func handle(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := req.URL.Query()
	bytevalue, _ := json.Marshal(query)
	fmt.Printf("%s\n", bytevalue)
	bufbody := bytes.NewBuffer(bytevalue)

	res, err := http.Post("http://localhost:8080/", "application/json", bufbody)
	if err != nil {
		fmt.Println("200 OK")
	}
	defer res.Body.Close()

	body, error := ioutil.ReadAll(res.Body)

	if error != nil {
		fmt.Println("200 OK")
	}

	fmt.Println(string(body))
	time.Sleep(time.Second * 1)

	w.Write(body)
}
