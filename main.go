package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var subPath *string
var port *string

func init(){
	subPath = flag.String("sub_path", "", "subPath")
	port = flag.String("port", "", "port")
	flag.Parse()
}

func getSub() []byte{
	file, err := os.Open(*subPath)
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
	http.ListenAndServe(*port, nil)
}