package main

import "fmt"

type Person interface {
	SaySomething()
}

type youngPerson struct {
	name string
	age  int
}

func (y youngPerson) SaySomething() {
	fmt.Printf("Hi, my name is %s and I am a young person", y.name)
}

type adultPerson struct {
	name string
	age  int
}

func (a adultPerson) SaySomething() {
	fmt.Printf("Hi, my name is %s and I am an adult person", a.name)
}

func NewPerson(name string, age int) Person {
	if age < 18 {
		return &youngPerson{name, age}
	}
	return &adultPerson{name, age}
}

func main() {
	p := NewPerson("John", 19)
	p.SaySomething()
}
