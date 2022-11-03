package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	// fmt.Println(args[0])
	// fmt.Println(args[1])

	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	fmt.Println("Hostname:", name)

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
