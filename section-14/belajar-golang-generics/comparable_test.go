package belajar_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func IsSameAny[T any](value1, value2 T) bool {
//	if value1 == value2 {
//		return true
//	}
//	return false
//}

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	}
	return false
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("100", "100"))
	assert.False(t, IsSame[int](100, 200))
}
