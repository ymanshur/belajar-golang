package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Yusuf"
	lastName = "Manshur"

	// return firstName, middleName, lastName
	return
}

func main() {
	first, middle, last := getFullName()
	// _, _, last := getFullName()

	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)
}
