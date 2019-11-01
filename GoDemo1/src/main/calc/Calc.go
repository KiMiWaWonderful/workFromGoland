package calc

import "fmt"

func init() {
	fmt.Println("this is calc's init")
}
func Add(a,b int) int {
	return a + b
}
func Minos(a,b int) int {
	return a - b
}

