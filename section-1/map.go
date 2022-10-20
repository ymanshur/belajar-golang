package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yusuf",
		"address": "Yogyakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "Programmer"
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Yusuf Manshur"
	book["wrong"] = "Ups"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}
