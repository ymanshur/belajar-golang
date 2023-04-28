package belajar_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type price int

type Number interface {
	~int | int8 | int16 | float32 | float64
}

func Min[T Number](value1, value2 T) T {
	if value1 < value2 {
		return value1
	}
	return value2
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	//assert.Equal(t, 100, Min[int64](100, 200))
	assert.Equal(t, int16(100), Min[int16](100, 200))
	assert.Equal(t, 100.0, Min[float64](100.0, 100.8))
	assert.Equal(t, price(27000), Min[price](price(30000), price(27000))) // should add approximation to base type ex: ~int
}

func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int16(100), Min(int16(100), int16(200)))
	assert.Equal(t, 100.0, Min(100.0, 100.8))
	assert.Equal(t, price(27000), Min(price(30000), price(27000)))
}
