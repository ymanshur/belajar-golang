package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulanagn ke", counter)
	// 	counter++
	// }

	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Perulanagn ke", counter)
	}

	slice := []string{"Muhammad", "Yusuf", "Manshur"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := map[string]string{
		"name":  "Yusuf",
		"title": "Programmer",
	}

	for key, value := range person {
		fmt.Println(key, value)
	}

}
