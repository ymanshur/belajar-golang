package belajar_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type DataA[T any] struct {
	Value T
}

func (d *DataA[T]) GetValue() T {
	return d.Value
}

func (d *DataA[T]) SetValue(value T) {
	d.Value = value
}

func TestGetterSetter(t *testing.T) {
	data1 := DataA[string]{
		Value: "1",
	}

	result := ChangeValue[string](&data1, "2")

	assert.Equal(t, result, data1.GetValue())
}
