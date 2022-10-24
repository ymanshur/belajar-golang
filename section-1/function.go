package main

import "fmt"

func sayHello() { // camel case
	fmt.Println("Hello World")
	fmt.Println("Hello World")
}

func main() {
	// sayHello()
	// sayHello()
	// sayHello()

	for i := 0; i < 5; i++ {
		sayHello()
	}
}
