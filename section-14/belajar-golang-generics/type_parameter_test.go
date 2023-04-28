package belajar_golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Print[T any](param T) T {
	fmt.Printf("%v %T\n", param, param)
	return param
}

func PrintAny(param interface{}) interface{} {
	fmt.Printf("%v %T\n", param, param)
	return param
}

func TestTypeParam(t *testing.T) {
	var result string = Print[string]("Yusuf")
	assert.Equal(t, "Yusuf", result)

	var resultNumber int = Print[int](100)
	assert.Equal(t, 100, resultNumber)

	var resultAny = PrintAny("100")
	assert.Equal(t, "100", resultAny)
}
