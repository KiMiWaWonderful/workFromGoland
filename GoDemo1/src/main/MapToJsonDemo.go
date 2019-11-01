package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{},4)
	m["company"] = "it"
	m["subjects"] = []string{"c++","go","c"}
	m["isok"] = false
	m["price"] = 6.66

	buf,err := json.MarshalIndent(m,""," ")

	if err != nil {
		fmt.Println("err = ",err)
		return
	}

	fmt.Println("buf = ",string(buf))
}