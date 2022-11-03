package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("salah")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(boolean)

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(number)

	value := strconv.FormatInt(1000, 10)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}
