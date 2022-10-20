package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPEko NoKTP = "13432423432423423"
	var marriedStatus Married = true
	fmt.Println(noKTPEko)
	fmt.Println(marriedStatus)
}
