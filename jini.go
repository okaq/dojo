// okaq bitmap emoji
// AQ <aq@okaq.com>
// 2019-04-07
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "jini.html"
)

var (
	// cache, concurrency-safe
	C map[string]string
	// counter, atomic
	U uint64
)

func motd() {
	fmt.Println(time.Now().String())
}

func JiniHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// increment counter
	http.ServeFile(w,r,INDEX)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// output visit count
}

func main() {
	motd()
	// init counter
	// init cache
	http.HandleFunc("/", JiniHandler)
	http.HandleFunc("/s", StatsHandler)
	http.ListenAndServe(":8080", nil)
}


