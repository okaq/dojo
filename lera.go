// okaq web server
// emoji bitmap test and layout engine
// AQ <aq@okaq.com>
// 2019-06-08
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "lera.html"
	NOJI = ""
)

func LeraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func NojiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,NOJI)
}

func motd() {
	fmt.Println("web serving on localhost:8080...")
	fmt.Printf("%s\n", time.Now().String())
}

func main() {
	motd()
	// rand()
	// cache()
	http.HandleFunc("/", LeraHandler)
	http.HandleFunc("/a", NojiHandler)
	http.ListenAndServe(":8080", nil)
}


