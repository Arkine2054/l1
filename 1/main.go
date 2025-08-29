package main

import "fmt"

type Human struct {
	name string
	game string
}

type Action struct {
	Human
}

func (h Human) sayHello() string {
	return fmt.Sprintf("Hello, my name is %s! I like %s\n", h.name, h.game)
}

func main() {
	person := Human{
		name: "Nikita",
		game: "basketball",
	}
	action := Action{
		Human: person,
	}

	fmt.Println(action.sayHello())
}
