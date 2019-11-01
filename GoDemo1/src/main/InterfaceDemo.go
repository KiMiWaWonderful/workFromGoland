package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call()  {
	fmt.Println("I am Nokia, i can call you")
}

type IPhone struct {

}

func (iPhone IPhone) call()  {
	fmt.Println("I am IPhone, i can call you")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

