package main

import (
	"encoding/json"
	"fmt"

	"github.com/ifreddyrondon/go-talks/santiago-feb2018/captures/capture"
)

func main() {
	// START OMIT
	in := []byte(`
		[
			{"lat": 1, "lng": 1, "date": "1989-12-26T06:01:00.00Z"},
			{"lat": 1, "lng": 2, "date": "1990-12-26T06:01:00.00Z"}
		]`,
	)
	// END OMIT
	var c capture.Captures
	if err := json.Unmarshal(in, &c); err != nil { // HL
		panic(err)
	}
	fmt.Println(c)
}
