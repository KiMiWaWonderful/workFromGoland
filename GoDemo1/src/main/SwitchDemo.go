package main

import "fmt"

type Stuu struct {
	name string
	id int
}

func main() {
	i := make([]interface{},3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Stuu{"mkie",22}

	for index, data := range i{
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]类型为int,内容为%d\n",index,value)
		case string:
			fmt.Printf("x[%d]类型为string,内容为%s\n",index,value)
		case Stuu:
			fmt.Printf("x[%d]类型为Stu,内容为%d,%s\n",index,value.id,value.name)
		}
	}
}
