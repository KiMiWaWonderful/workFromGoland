package main

import "fmt"

type FuncType func(int,int) int

func Add(a,b int) int {
	return a+b
}
func Calc(a,b int,fTest FuncType)(result int){
	fmt.Println("Calc")
	result = fTest(a,b)
	return result
}

func main() {
	a := Calc(1,1,Add)
	fmt.Println("a=",a)
}
