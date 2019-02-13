// okaq web serve
// unicode glyph bitmap sample 
// server base64 json encode
// AQ <aq@okaq.com>
// 2019-02-14
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "doho.html"
	DATA = "dohod/"
)

func DohoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Printf("%s\n", time.Now().String())
}

func main() {
	motd()
	http.HandleFunc("/", DohoHandler)
	http.ListenAndServe(":8080", nil)
}


