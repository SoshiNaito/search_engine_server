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
type Res []struct {
	URL string `json:"url"`
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
		// return
	}
	defer res.Body.Close()

	var r Res

	body, error := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	error = json.Unmarshal(body, &r)
	if error != nil {
		log.Fatal(error)
	}

	// sb := string(body)
	// fmt.Fprintf(w, "%v\n", body)
	w.Write(body)
}
