package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// START OMIT
	in := []byte(`{"hello": "world"}`)
	var result map[string]interface{}
	if err := json.Unmarshal(in, &result); err != nil { // HL
		panic(err)
	}
	fmt.Println(result)
	// END OMIT
}
