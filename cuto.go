// okaq nano game
// hanabi for AI
// AQ <aq@okaq.com>
// 2019-02-06
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "cuto.html"
)

func CutoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func motd() {
	fmt.Println("okaq web serv on localhost:8080")
	fmt.Printf("start: %s\n", time.Now().String())
}

func main() {
	motd()
	http.HandleFunc("/", CutoHandler)
	http.ListenAndServe(":8080", nil)
}


