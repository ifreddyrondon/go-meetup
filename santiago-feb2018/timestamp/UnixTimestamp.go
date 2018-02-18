package main

import (
	"encoding/json"
	"fmt"

	"github.com/ifreddyrondon/go-talks/santiago-feb2018/timestamp/timeutils"
)

func main() {
	// START OMIT
	in := []byte(`{"date": 630655260}`) // HL
	// END OMIT
	var t timeutils.Timestamp
	if err := json.Unmarshal(in, &t); err != nil { // HL
		panic(err)
	}
	fmt.Println(t)
}
