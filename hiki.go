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
}

func HikiHandler() {
}

func main() {
	motd()
	http.HandleFunc("/", HikiHandler)
	http.ListenAndServe(":8080", nil)
}


