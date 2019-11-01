package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a,b)

	fmt.Printf("最大值是： %d\n",ret)

	c, d := swap("Google","Runoob")
	fmt.Print(c,d)
}
func max(num1, num2 int) int{
	var result int

	if (num1 > num2){
		result = num1
	}else {
		result = num2
	}
	return result
}

func swap(x,y string)(string, string)  {
	return y, x
}

func swap2(x, y int)int  {
	var temp int

	temp = x
	x = y
	y = temp

	return temp
}

func swap3(x,y int)int  {
	var temp int

	temp = *x
	*x = *y
	*y = temp

	return temp
}
