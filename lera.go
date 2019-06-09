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
	NOJI = "noto_emoji_2.json"
	GOLF = "golf_1.json"
)

func LeraHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func NojiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,NOJI)
}

func GolfHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,GOLF)
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
	http.HandleFunc("/b", GolfHandler)
	http.ListenAndServe(":8080", nil)
}


