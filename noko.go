// favicon generate and encode
// AQ <aq@okaq.com>
// 2018-12-25
// merry xmas!
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

const (
	PATH = "nokoa/"
	PNG = "favicon_"
	PNG_E = ".png"
	JSON = "dataUri_"
	JSON_E = ".json"
)

var (
	// png path
	P string
	// png file
	Pn *os.File
	// json path
	J string
	// json file
	Js *os.File
	// time
	T time.Time
	// pixels
	I *image.RGBA
)

func Files() {
	fmt.Println("opening files to save on disk")
	t := T.UnixNano()
	p := fmt.Sprintf("%s%s%d%s", PATH, PNG, t, PNG_E)
	j := fmt.Sprintf("%s%s%d%s", PATH, JSON, t, JSON_E)
	fmt.Println(p, j)
	var err error
	Pn, err = os.Create(p)
	if err != nil {
		panic(err)
	}
	Js, err = os.Create(j)
	if err != nil {
		panic(err)
	}
}

func Pixels() {
	r := image.Rect(0,0,32,32)
	I = image.NewRGBA(r)
}

func Solid() {
	// generate a solid color favicon
	// c := color.RGBA{128,128,128,255}
	// dojo aomi, basic off white
	c := color.RGBA{216,216,216,255}
	dx := I.Bounds().Dx()
	dy := I.Bounds().Dy()
	b := dx * dy
	for i := 0; i < b; i++ {
		x := i % dx
		y := i / dx
		I.SetRGBA(x,y,c)
	}
}

func Square() {
	// two color square within a square
	c0 := color.RGBA{250,250,250,255}
	c1 := color.RGBA{10,10,10,255}
	dx := I.Bounds().Dx()
	dx2 := dx / 4
	dx3 := 3 * dx / 4
	dy := I.Bounds().Dy()
	dy2 := dy / 4
	dy3 := 3 * dy / 4
	b := dx * dy
	for i := 0; i < b; i++ {
		x := i % dx
		y := i / dx
		if x > dx2 && x <= dx3 && y > dy2 && y <= dy3 {
			I.SetRGBA(x,y,c0)
		} else {
			I.SetRGBA(x,y,c1)
		}
	}
}

func Write() {
	// png file
	err := png.Encode(Pn, I)
	if err != nil {
		fmt.Println(err)
	}
	// png data
	f, err := Pn.Stat()
	if err != nil {
		fmt.Println(err)
	}
	d := make([]byte, f.Size())
	_, err = Pn.Seek(0,0)
	if err != nil {
		fmt.Println(err)
	}
	_, err = Pn.Read(d)
	if err != nil {
		fmt.Println(err)
	}
	// json encode
	s := base64.StdEncoding.EncodeToString(d)
	fmt.Println(f,d,s)
	uri := fmt.Sprintf("data:image/png;base64,%s", s)
	fmt.Println(uri)
	fav := struct { DataURI string } { uri }
	fmt.Println(fav)
	json, err := json.Marshal(fav)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
	_, err = Js.Write(json)
	if err != nil {
		fmt.Println(err)
	}
	g, err := Js.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("png and json files sccessfully encoded and written to disk")
	fmt.Printf("png size: %d bytes\n", f.Size())
	fmt.Printf("json size: %d bytes\n", g.Size())
}

func main() {
	T = time.Now()
	fmt.Printf("okaq favicon generator start: %s\n", time.Now().String())
	Files()
	defer Pn.Close()
	defer Js.Close()
	Pixels()
	// generator
	Solid()
	// Square()
	Write()
}

// abstract concentric box function
// input a list of colors
// divide space and set layers


