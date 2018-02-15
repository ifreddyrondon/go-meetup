package main

import (
	"encoding/json"
	"fmt"

	"github.com/ifreddyrondon/go-talks/santiago-feb2018/structs/point"
)

func main() {
	// START OMIT
	in := []byte(`{"lat": 1.1, "lng": 2.2}`) // HL

	var p point.Point
	if err := json.Unmarshal(in, &p); err != nil { // HL
		panic(err)
	}
	fmt.Println(p)
	// END OMIT
}
