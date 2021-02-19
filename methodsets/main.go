package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hello, I am", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Sherlock",
		last:  "Holmes",
		age:   32,
	}
	//saySomething(p1)
	saySomething(&p1)
}

//saySomething(p1) is not allowed. This is because speak method has a pointer receiver.
//cannot use p1 (variable of type person) as human value
//in argument to saySomething: missing method speak (speak has pointer receiver)
