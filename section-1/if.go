package main

import "fmt"

func main() {
	var name = "Joko"

	if name == "Yusuf" {
		fmt.Println("Hello Yusuf")
	} else if name == "Manshur" {
		fmt.Println("Hello Manshur")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	// var length = len(name)
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
