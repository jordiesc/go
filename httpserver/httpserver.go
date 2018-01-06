/**
* http server to mok
 */
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	port := flag.Int("port", 8080, "port per defecte")
	cors = flag.Bool("cors", true, "descativa CORS per defecte esta activa")
	flag.Parse()

	sport := strconv.Itoa(*port)
	sport = ":" + sport

	fmt.Print("servidor pel port " + sport)

	http.HandleFunc("/", handler)
	http.ListenAndServe(sport, nil)

}

var cors *bool

func handler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.RequestURI, "html") {
		w.Header().Add("Content-Type", "text/html; charset=UTF-8")
		w.Header().Add("Connection", "keep-alive")

	} else {
		w.Header().Add("content-type", "application/json")
	}
	if *cors {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	}

	buf := new(bytes.Buffer)
	r.Write(buf)
	fmt.Printf(buf.String())
	fmt.Printf(time.Now().String() + "\n")
	fmt.Printf("llegint fitxer %d" + r.RequestURI)
	fmt.Fprintf(w, readFile(r.RequestURI))
}

/*
* read a file or a json file without extension
 */
func readFile(filename string) string {
	var fi []byte
	fi, err := ioutil.ReadFile("." + filename)
	if err != nil {
		fmt.Println(err)
	}
	return (string(fi))
}
