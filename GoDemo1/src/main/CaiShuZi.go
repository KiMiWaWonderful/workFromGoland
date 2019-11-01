package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int)  {
	rand.Seed(time.Now().UnixNano())

	var num int
	for{
		num = rand.Intn(10000)
		if num >= 1000{
			break
		}
	}
	*p = num
}

func main() {
	var randNum  int
	CreateNum(&randNum)
	fmt.Println("randNum:",randNum)
}
