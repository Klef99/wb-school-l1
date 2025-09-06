package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h *Human) Greetings() string {
	return fmt.Sprintf("Hi, my name is %s! I am %d years old!", h.Name, h.Age)
}

func NewHuman(name string, age int) *Human {
	return &Human{name, age}
}

type Action struct {
	Human
}

func (a *Action) Do() string {
	return fmt.Sprintf("Do work")
}

func NewAction(human Human) *Action {
	return &Action{human}
}

const (
	name = "Egor"
	age  = 20
)

func main() {
	h := NewHuman(name, age)
	action := NewAction(*h)

	fmt.Println(action.Greetings())
	fmt.Println(action.Do())
}
