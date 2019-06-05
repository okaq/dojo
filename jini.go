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
	FONTS = "fonts/NotoEmoji-Regular.ttf"
)

var (
	// cache, concurrency-safe
	C map[string]string
	// counter, atomic
	U uint64
	// accumulator
	// G <- int
)

func motd() {
	fmt.Println(time.Now().String())
}

func counter() {
	// launch goroutine
	// and for loop reciever
}

func cache() {
	// key is user id
	// value is key into db
	// store pid, time, conn data
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

func PidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// two round trips
	// one on initial load
	// two on first player interaction
}

func ExtHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// admin utilities
	// monitor peer net, alter state
}

func FontHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,FONTS)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// store the sampled bit array
	// create bytes buffer
	b0 := new(bytes.Buffer)
	b0.ReadFrom(r.Body)
	fmt.Println(b0.Bytes())
	fmt.Sprintf("Recieved %d bytes from sampler", b0.Len())
	// parse json
	j0, err := json.Marshal(b0.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	// write to disk
}

func main() {
	motd()
	// init counter
	counter()
	// init cache
	cache()
	http.HandleFunc("/", JiniHandler)
	http.HandleFunc("/s", StatsHandler)
	http.HandleFunc("/a", PidHandler)
	http.HandleFunc("/x", ExtHandler)
	http.HandleFunc("/e", FontHandler)
	http.HandleFunc("/d", SaveHandler)
	http.ListenAndServe(":8080", nil)
}

// unicode code point ranges for emoji glyphs
// '\u1f994' = 'hedgehog'
// '\u1f987' = 'bat'


