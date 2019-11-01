package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(strings.Contains("acb","a"))

//	s := make([]string,{"abc","aaa","mike"})
	s := []string{"abc","hello","world"}
	buf := strings.Join(s,",")
	fmt.Println(buf)
	fmt.Println(strings.Index("hellowrold","g"))
	fmt.Println(strings.Repeat("go",3))
	fmt.Println(strings.Split("hello,string,int",","))
}
