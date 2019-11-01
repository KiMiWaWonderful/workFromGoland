package main

import "unsafe"

const (
	x = "abc"
	y = len(x)
	z = unsafe.Sizeof(x)
)

func main() {
	println(x,y,z)
}
