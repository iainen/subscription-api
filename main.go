package main

import (
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



func getSub() string{
	file, err := os.Open(*subPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content)
}

func SubHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getSub())
}

func main () {
	http.HandleFunc("/", SubHandler)
	http.ListenAndServe(*port, nil)
}