package belajar_golang_generics

import (
	"fmt"
	"testing"
)

func MultipleParam[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Printf("%v %T\n", param1, param1)
	fmt.Printf("%v %T\n", param2, param2)
}

func TestMultipleParam(t *testing.T) {
	MultipleParam[string, int]("100", 100)
	MultipleParam[int, string](100, "100")
}
