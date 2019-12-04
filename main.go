package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Res struct
type Res struct {
	Ts []string
}

func main() {
	http.HandleFunc("/", jsonhandle)

	http.ListenAndServe(":9080", nil)

}

func jsonhandle(w http.ResponseWriter, req *http.Request) {
	m := req.URL.Query()
	
	bytevalue, _ := json.Marshal(m)

	fmt.Printf("%s\n", bytevalue)

	res, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer(bytevalue))
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	fmt.Printf("[status] %d\n", res.StatusCode)

	defer res.Body.Close()
	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}
	sb := string(body)
	fmt.Printf("[body] %v\n", sb)

}
