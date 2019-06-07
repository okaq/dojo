// okaq bitmap emoji
// AQ <aq@okaq.com>
// 2019-04-07
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	INDEX = "jini.html"
	FONTS = "fonts/NotoEmoji-Regular.ttf"
	NOTO = "fonts/noto_emoji_1.json"
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
	fmt.Println(r.Body)
	b0 := new(bytes.Buffer)
	b0.ReadFrom(r.Body)
	fmt.Println(b0.Bytes())
	fmt.Printf("Recieved %d bytes from sampler", b0.Len())
	// parse json
	j0, err := json.Marshal(b0.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j0))
	// unmarshal into map[string][]byte
	// marshal to {"unicode":base64}
	var d0 map[string][]byte
	err = json.Unmarshal(b0.Bytes(), &d0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d0)
	j1, err := json.Marshal(d0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(j1)
	// write to disk
	// ioutil.WriteFile("noto_emoji.json", j0, 0666)
	ioutil.WriteFile(NOTO, j1, 0666)
	s0 := fmt.Sprintf("{bytes:%d}", len(j0))
	b1 := []byte(s0)
	w.Write(b1)
}

func NotoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,NOTO)
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
	http.HandleFunc("/n", NotoHandler)
	http.ListenAndServe(":8080", nil)
}

// unicode code point ranges for emoji glyphs
// '\u1f994' = 'hedgehog'
// '\u1f987' = 'bat'


