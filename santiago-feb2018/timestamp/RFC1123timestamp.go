package main

import (
	"encoding/json"
	"fmt"

	"github.com/ifreddyrondon/go-talks/santiago-feb2018/timestamp/timeutils"
)

func main() {
	// START OMIT
	in := []byte(`{"timestamp": "Tue, 26 Dec 1989 06:01:00 UTC"}`) // HL
	// END OMIT
	var t timeutils.Timestamp
	if err := json.Unmarshal(in, &t); err != nil { // HL
		panic(err)
	}
	fmt.Println(t)
}
