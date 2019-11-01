package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}

type Studentt struct {
	Person
	id int
	addr string
}

func (temp Person)PrintInfo()  {
	fmt.Println(temp)
}

func (temp Studentt)PrintInfo()  {
	fmt.Println(temp)
}

func (temp *Person)SetInfo(n string,s byte,a int)  {
	temp.name = n
	temp.sex = s
	temp.age = a
}

func main() {
	s1 := Studentt{Person{"mike",'m',12},1,"beijing"}
	s1.Person.PrintInfo()
	sFunc := s1.PrintInfo
	sFunc()
	//	//fmt.Printf("s1 = %+v\n",s1)
	//	//fmt.Println(s1)
   //p := Person{"mike",'x',10}
   //p.PrintInfo()

   //var p Person
	//(&p).SetInfo("mke",'s',100)
   //p.PrintInfo()

}
