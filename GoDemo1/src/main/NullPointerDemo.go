package main

import "fmt"

func main() {
	var ptr *int

	if(ptr == nil){
		fmt.Println("空")
		fmt.Printf("ptr的值为：%x\n",ptr);
	}else {
		fmt.Printf("ptr的值为：%x\n",ptr);
	}
}
