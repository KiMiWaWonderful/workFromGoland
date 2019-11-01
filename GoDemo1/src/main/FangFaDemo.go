package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle)getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var c1  Circle
	c1.radius = 10.00
	fmt.Println("圆的面积=",c1.getArea())
}