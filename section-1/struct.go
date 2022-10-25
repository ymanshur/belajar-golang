package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	isMarried     bool
}

func main() {
	var yusuf Customer
	yusuf.Name = "Yusuf"
	yusuf.Address = "Yogya"
	yusuf.Age = 25

	fmt.Println(yusuf.Name)
	fmt.Println(yusuf.Address)
	fmt.Println(yusuf.Age)

	var manshur = Customer{
		Name:    "Manshur",
		Address: "Sleman",
		Age:     24,
	}
	fmt.Println(manshur)

	muhammad := Customer{"Muhammad", "Wates", 23}
	fmt.Println(muhammad)
}
