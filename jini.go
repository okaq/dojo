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

)

func motd() {
	fmt.Println(time.Now().String())
}

func JiniHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", JiniHandler)
	http.ListenAndServe(":8080", nil)
}


