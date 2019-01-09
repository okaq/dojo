// okaq web server
// AQ <aq@okaq.com
// 2019-01-10
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "aomi.html"
)

func motd() {
	fmt.Println("okaq web start on localhost:8080")
	fmt.Printf("%s\n", time.Now().String())
}

func AomiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", AomiHandler)
	http.ListenAndServe(":8080", nil)
}


