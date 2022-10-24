package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Kenalan yuk!"
	}
	result := "Hello " + name
	return result

	// fmt.Println(result)
}

func main() {
	name := "Yusuf"
	hello := getHello(name)
	fmt.Println(hello)

	hello2 := getHello("")
	fmt.Println(hello2)
}
