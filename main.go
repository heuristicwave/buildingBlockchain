package main

import "fmt"

type person struct {
	name string
	age  int
}

func (instance person) sayHi() {
	fmt.Printf("Hi, My name is %s and %d", instance.name, instance.age)
}

func main() {
	my := person{"heuri", 7}
	my.sayHi()
}
