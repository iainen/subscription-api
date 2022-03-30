package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getSub() []byte{
	file, err := os.Open("config/sub.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return content
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	sEnc:= base64.StdEncoding.EncodeToString(getSub())
	fmt.Fprintf(w, sEnc)
}

func main () {
	http.HandleFunc("/sub", SubHandler)
	http.ListenAndServe(":8080", nil)
}
