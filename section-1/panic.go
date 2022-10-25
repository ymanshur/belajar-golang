package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")

	errMessage := recover()
	if errMessage != nil {
		fmt.Println("Terjadi error:", errMessage)
	}

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

	fmt.Println("Yusuf")
}
