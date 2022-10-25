package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	isMarried     bool
}

// func sayHello(customer Customer, name string) {
// 	fmt.Println("Hello", name, "my name is", customer.Name)
// }

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var yusuf Customer
	yusuf.Name = "Yusuf"
	yusuf.Address = "Yogya"
	yusuf.Age = 25

	fmt.Println(yusuf.Name)
	fmt.Println(yusuf.Address)
	fmt.Println(yusuf.Age)

	// sayHello(yusuf, "Manshur")
	yusuf.sayHello("Manshur")

	// var manshur = Customer{
	// 	Name:    "Manshur",
	// 	Address: "Sleman",
	// 	Age:     24,
	// }
	// fmt.Println(manshur)

	// muhammad := Customer{"Muhammad", "Wates", 23}
	// fmt.Println(muhammad)
}
