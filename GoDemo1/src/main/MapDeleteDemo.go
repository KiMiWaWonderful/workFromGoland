package main

import "fmt"

func main() {
	countryCaptialMap := map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New delhi"}
	fmt.Println("原始地图 ")

	for country := range countryCaptialMap{
		fmt.Println(country,"首都是：",countryCaptialMap[country])
	}

	delete(countryCaptialMap,"France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除元素后地图")

	for country := range countryCaptialMap{
		fmt.Println(country,"首都是",countryCaptialMap[country])
	}
}
