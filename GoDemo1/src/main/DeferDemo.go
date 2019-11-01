package main

import "fmt"

func test(x int) int {
	result := 100/x
	return result
}
func main() {
	defer fmt.Println("aaaaa")
	defer fmt.Println("bbbb")
}
