package belajar_golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Yusuf",
		Second: "Manshur",
	}

	data2 := Data[int]{
		First:  1,
		Second: 2,
	}

	fmt.Println(data)
	fmt.Println(data2)

	assert.Equal(t, "Hello Yusuf", data.SayHello("Yusuf"))
	assert.Equal(t, "Manshur", data.ChangeFirst("Manshur"))
}
