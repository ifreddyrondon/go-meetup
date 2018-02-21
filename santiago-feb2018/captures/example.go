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
			{"lat": 40, "lng": 20, "date": "1990-12-26T06:01:00.00Z"},
			{"lat": 50, "lng": 100, "date": "1991-12-26T06:01:00.00Z"}
		]`,
	)
	var c capture.Captures
	if err := json.Unmarshal(in, &c); err != nil {
		panic(err)
	}
	// END OMIT

	response, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(response))
}
