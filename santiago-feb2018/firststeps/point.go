package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// START OMIT
	type Point struct { // HL
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	in := []byte(`{"lat": 1.1, "lng": 2.2}`)

	var point Point                                    // HL
	if err := json.Unmarshal(in, &point); err != nil { // HL
		panic(err)
	}
	fmt.Println(point)
	// END OMIT
}
