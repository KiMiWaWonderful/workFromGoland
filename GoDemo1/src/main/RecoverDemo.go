package main

import "fmt"

func testa()  {
	fmt.Println("aaaaaa")
}

func testb(x int)  {
	defer func() {
		//recover()
		if err := recover(); err != nil{
			fmt.Println(recover())
		}
	}()

	var a [10]int
	a[x] = 111
}

func testc()  {
	fmt.Println("cccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
