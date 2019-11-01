package main

import (
	"fmt"
	"main/calc"
)

func init() {
	fmt.Println("this is main's init")
}
func main() {
	a := calc.Add(10,20)
	fmt.Println("a=",a)
}
