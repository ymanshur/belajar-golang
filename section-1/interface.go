package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups() interface{} {
	return 1
	return true
	return "Ups"
}

func main() {
	var yusuf Person
	yusuf.Name = "Yusuf"

	sayHello(yusuf)

	var cat Animal
	cat.Name = "Abu"

	sayHello(cat)

	var ups interface{} = Ups()
	fmt.Println(ups)
}
