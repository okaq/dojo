// okaq web serve
// wave function collapse
// AQ <aq@okaq.com>
// 2019-02-23
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "erai.html"
)

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Printf("%s\n", time.Now().String())
}

func meta() {
	// init log files
	fmt.Println("generating log files...")
}

func EraiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	meta()
	http.HandleFunc("/", EraiHandler)
	http.ListenAndServe(":8080", nil)
}


