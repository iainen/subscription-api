package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getSub(token string) []byte {
	path := fmt.Sprintf("config/%s.txt", token)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return content
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	sEnc := base64.StdEncoding.EncodeToString(getSub(token))
	fmt.Fprintf(w, sEnc)
}

func main() {
	http.HandleFunc("/sub", SubHandler)
	http.ListenAndServe(":8080", nil)
}
