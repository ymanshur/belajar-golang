package main

import "fmt"

func main() {
	name := "Yusuf"
	counter := 0

	increment := func() {
		// name = "Manshur"
		name := "Manshur"
		fmt.Println("Increment")
		counter++

		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
