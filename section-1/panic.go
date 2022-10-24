package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(err bool) {
	defer endApp()

	if err {
		panic("ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
