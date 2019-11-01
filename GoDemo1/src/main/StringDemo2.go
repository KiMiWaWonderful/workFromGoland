package main

import (
	"fmt"
	"strconv"
)

func main() {
	sli := make([]byte,0,1024)
	sli = strconv.AppendBool(sli,false)
	//sli = strconv.AppendFloat(sli,)
	sli = strconv.AppendInt(sli,100,10)
	fmt.Println(string(sli))

}
