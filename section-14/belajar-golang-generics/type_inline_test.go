package belajar_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Max[T interface {
	~int | int8 | int16 | float32 | float64
}](first, second T) T {
	if first > second {
		return first
	}
	return second
}

func TestMax(t *testing.T) {
	assert.Equal(t, 100, Max(50, 100))
	assert.Equal(t, 100.6, Max(100.0, 100.6))
}

func GetFirst[T []E, E any](slice T) E {
	return slice[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{"Yusuf", "Manshur"}
	numbers := []int{1, 2}

	first := GetFirst(names)
	first2 := GetFirst(numbers)

	assert.Equal(t, "Yusuf", first)
	assert.Equal(t, 1, first2)
}
