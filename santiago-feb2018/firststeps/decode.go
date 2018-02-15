package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// START OMIT
	in := strings.NewReader(`{"hello": "world"}`)
	var result map[string]interface{}
	json.NewDecoder(in).Decode(&result) // HL
	fmt.Println(result)
	// END OMIT
}
