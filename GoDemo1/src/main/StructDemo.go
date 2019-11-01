package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func ChangeStruct(s *Student)  {
	s.id = 10
	fmt.Println("in struct",s)
}
func main() {
	//var s1 *Student = &Student{id:1}
	//fmt.Println(*s1)

	var s1 Student
	s1.addr = "beijing"
		s1.name = "mike"
	ChangeStruct(&s1)
		fmt.Println(s1)
	//	//fmt.Println(s1)
	//var p1 *Student
	//p1 = &s1
	//
	//fmt.Println(p1.name)
	//fmt.Println(p1)
	//
	//p2 := new(Student)
	//fmt.Println(p2)
}
