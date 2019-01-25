// okaq nano game
// turn based strategy
// move and confront 
// AQ <aq@okaq.com>
// 2019-01-26
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "bomi.html"
)

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Printf("timestamp: %s\n", time.Now().String())
}

func BomiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", BomiHandler)
	http.ListenAndServe(":8080", nil)
}


