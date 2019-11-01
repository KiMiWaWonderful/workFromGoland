package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string
	Subjects []string
	IsOk bool
	Price float64
}

func main() {
	s := IT{"acb",[]string{"c++","go","c"},true,66.6}

	buf,err := json.MarshalIndent(s,""," ")

	if err != nil {
		fmt.Println("err = ",err)
		return
	}

	fmt.Println("buf = ",string(buf))
}
