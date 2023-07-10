package main

import "fmt"

// Структура Human с полем name и методами sing() и walk()
type Human struct {
	name string
}

func (h Human) sing() {
	fmt.Printf("%s поет: ла-ла-ла\n", h.name)
}

func (h Human) walk() {
	fmt.Printf("%s идет: топ-топ\n", h.name)
}

// Структура Action с полем friend, в которую встроена Human, и методом beFriends()
type Action struct {
	Human
	friend string
}

func (a Action) beFriends() {
	fmt.Printf("%s и %s теперь друзья\n", a.name, a.friend)
}

func main() {
	// Создаем экземпляр структуры Human
	human := Human{
		name: "Alexander",
	}
	// Вызов методов sing() и walk() у структуры Human
	human.sing()
	human.walk()

	// Создаем экземпляр структуры Action
	action := Action{
		Human:  Human{name: "Alexander"},
		friend: "Maria",
	}
	// Вызов метода beFriends(), а также методов встроенной стрктуры Human, у структуры Action
	action.sing()
	action.walk()
	action.beFriends()
}
