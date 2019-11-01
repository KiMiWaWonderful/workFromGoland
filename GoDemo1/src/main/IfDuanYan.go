package main

import "fmt"

type Stu struct {
	name string
	id int
}

func main() {
	i := make([]interface{},3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Stu{"mkie",22}

	for index, data := range i{
		if value, ok := data.(int); ok == true{
			fmt.Printf("x[%d]类型为int,内容为%d\n",index,value)
		}else if value, ok := data.(string); ok == true{
			fmt.Printf("x[%d]类型为string,内容为%s\n",index,value)
		}else if value, ok := data.(Stu); ok == true{
			fmt.Printf("x[%d]类型为Stu,内容为%d,%s\n",index,value.id,value.name)
		}
	}
}
