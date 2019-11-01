package main

import "fmt"

type MyInt int

const Pi  float64 = 3.14159
const zero  = 0.0
const u,v float32  = 0,3
//多重赋值
const a,b,c  = 3,4,"foo"
//用作枚举
const(
	size int64 = 1024
	eof = -1
)
func add(args...int)int  {
	sum := 0;
	for _,arg := range args{
		sum += arg
	}
	return sum
}

//func f(a,b int)(int,error){
//
//}

type Integer int

func (a Integer) Add(b Integer)  Integer {
	return a + b
}

func (a *Integer) Addd(b Integer)  Integer {
	return *a + b
}
func main() {
	x := 1
	{
		x := 2
		fmt.Print(x)
	}
	fmt.Print(x)
	//var s[]int
	//s = append(s,1)
	//var m map[string]int
	//m = make(map[string]int)
	//m["one"] = 1
	//var a Integer = 1
	//var b Integer = 2
	//var i interface{} = a
	//sum := i.(Integer).Add(b)
	//fmt.Println(sum)
	//i := 1
	////j = i++
	//i--
	//fmt.Println(i)
	//numbers := []int{0,1,2,3,4,5,6,7,8}
	//s := make([]int,5,10)
	//ss := []int{
	//	1,2,3,
	//2,3,4}
	//sss := numbers[:]

	//var x interface{}= nil
	//if(true){
	//	defer fmt.Printf("1")
	//}else {
	//	defer fmt.Printf("2")
	//}
	//fmt.Printf("3")
	//var b  = bool(1)
	//fmt.Println(Pi)
	//fmt.Println(zero)
	//fmt.Println(size)
	//fmt.Println(eof)
	//var i int
	// i = add([]int{1,2,3}...)
	// fmt.Println(i)

	//var i int = 1
	//var j = MyInt(i)
	//fmt.Println(j)
}
