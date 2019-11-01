package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//func (i int,j int) {
//	fmt.Println("......")
//}(10,20)
func main() {
	//nextNumber := getSequence()
	//
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//
	//nextNumber1 := getSequence()
	//
	//fmt.Println(nextNumber1())
	//fmt.Println(nextNumber1())

	func(i,j int) {
		fmt.Println(",,,,")
	}(1,2)
}