package main

import "fmt"

type Humaner interface {
	sayhi()
}

type Personer interface {
	Humaner
	sing(lrc string)
}

type Studenttt struct {
	name string
	id int
}

type Teacher struct {
	addr string
	group string
}

func (tmp *Studenttt) sayhi()  {
	fmt.Printf("Student[%s,%d] sayhi\n",tmp.name,tmp.id)
}

func (tmp *Studenttt) sing(lrc string)  {
	fmt.Printf("Student[%s,%d] sing\n",tmp.name,tmp.id)
}

func (tmp *Teacher)sayhi()  {
	fmt.Printf("Teacher[%s,%s] sayhi\n",tmp.addr,tmp.group)
}

func (tmp *Teacher) sing(lrc string)  {
	fmt.Printf("Teacher[%s,%s] sing\n",tmp.addr,tmp.group)
}

func WhoSayHi(i Humaner)  {
	i.sayhi()
}
func main() {
	//var i Humaner
	//
	//s := &Studenttt{"Mike",666}
	//t := &Teacher{"beijing","yuwen"}
	////i = s
	////i.sayhi()
   //
	//x := make([]Humaner,2)
	//x[0] = s
	//x[1] = t
   //
   //for _, i := range x{
   //	i.sayhi()
   //}

   var i Personer
	s := &Studenttt{"Mike",666}
	i = s
	i.sayhi()
	i.sing("....")
	//t := &Teacher{"beijing","yuwen"}

}