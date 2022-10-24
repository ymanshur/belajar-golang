package main

import "fmt"

func getFullName() (string, string) {
	return "Yusuf", "Manshur"
}

func main() {
	// firstName, lastName := getFullName()
	_, lastName := getFullName()

	// fmt.Println(firstName)
	fmt.Println(lastName)
}
