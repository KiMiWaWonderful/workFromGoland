package main

import "fmt"

func main() {
	var b int = 15
	var a int

	numbers := [6]int{1,2,3,5}

	for a := 0;a < 10; a++ {
		fmt.Printf("a的值为：%d\n",a)
	}

	for a < b {
		a ++
		fmt.Printf("a的值为：%d\n",a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位x的值 = %d\n",i,x)
	}
}
