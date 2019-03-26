// okaq web serve
// koi koi game server
// AQ <aq@okaq.com>
// 2019-03-21
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "hiki.html"
)

var (
	// cache, rng
)

func motd() {
	fmt.Println("web serving on localhost:8080")
	fmt.Printf("%s\n", time.Now().String())
}

func HikiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func StatHandler(w http>responseWriter, r *http.Request) {
	fmt.Println(r)
	// pretty print state cache to writer
}

func main() {
	motd()
	http.HandleFunc("/", HikiHandler)
	http.HandleFunc("/s", StatHandler)
	http.ListenAndServe(":8080", nil)
}


