package main

import (
	"errors"
	"fmt"
)

func MyDiv(a,b int)(result int,err error)  {
	err = nil
	if b==0{
		err = errors.New("分母不能为0")
	}else {
		result = a/b
	}
	return
}

func main() {
	//err1 := fmt.Errorf("%s","this is normal errl")
	//fmt.Println("err1=",err1)

	result,err := MyDiv(10,0)
	if(err != nil){
		fmt.Println("err = ",err)
	}else {
		fmt.Println("result = ",result)
	}
}
