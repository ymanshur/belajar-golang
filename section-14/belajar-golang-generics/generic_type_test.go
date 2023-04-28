package belajar_golang_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBagAny[T any](bag []T) {
	for _, b := range bag {
		fmt.Println(b)
	}
}

//func PrintBag(bag Bag[string])  {
//	for _, b := range bag {
//		fmt.Println(b)
//	}
//}

//func PrintBag(bag Bag[int])  {
//	for _, b := range bag {
//		fmt.Println(b)
//	}
//}

func PrintBag[T any](bag Bag[T]) {
	for _, b := range bag {
		fmt.Println(b)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Muhammad", "Yusuf", "Manshur"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3}
	PrintBag(numbers)
}
