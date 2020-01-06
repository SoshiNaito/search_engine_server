package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Ranking struct
type Ranking struct {
	Rank int
	Res
}

//Res struct
type Res struct {
	ID int
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)

}

func handle(w http.ResponseWriter, req *http.Request) {

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
		return
	}
	defer res.Body.Close()

	var r Ranking

	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		log.Fatal(error)
	}

	error = json.Unmarshal(body, &r)
	if error != nil {
		log.Fatal(error)
	}
	sb := string(body)
	fmt.Printf("[body] %v\n", sb)
	//fmt.Printf("%v\n", r)
	newbytebody := bytes.NewBuffer(body)
	secondres, err := http.Post("http://localhost:9080/", "application/json", newbytebody)
	if err != nil {
		fmt.Println("Request error :", err)
		return
	}
	defer secondres.Body.Close()

	var url interface{}

	secondbody, error := ioutil.ReadAll(secondres.Body)
	if error != nil {
		log.Fatal(error)
	}
	error = json.Unmarshal(secondbody, &url)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%v\n", url)
}
