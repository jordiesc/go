package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"net/http"
	"bytes"
)

func main(){

	fmt.Printf(time.Now().String()+"\n")	
	fmt.Println(readFile("json"))

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler (w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	r.Write(buf)
	fmt.Printf(buf.String())

	w.Header().Add("content-type","application/json")
	fmt.Fprintf(w,readFile("json"))
	//w.Write([]byte(readFile("json")))
	
}
/*
* read a file or a json file without extension
*/
func readFile( filename string) string {
	var fi []byte
	fi,err := ioutil.ReadFile(filename)
	check(err)
	return (string(fi))
}



func check(e error) {
    if e != nil {
        panic(e)
    }
}