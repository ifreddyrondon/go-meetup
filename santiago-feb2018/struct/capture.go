package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// START OMIT
	type Point struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	in := []byte(`{"lat": 1.1, "lng": 2.2}`)

	var point Point
	if err := json.Unmarshal(in, &point); err != nil {
		panic(err)
	}
	fmt.Println(point)
	// END OMIT
}
