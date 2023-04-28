package belajar_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"testing"
)

func ExpMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestExpMin(t *testing.T) {
	assert.Equal(t, 100, ExpMin(100, 200))
	assert.Equal(t, int16(100), ExpMin(int16(100), int16(200)))
	assert.Equal(t, 100.0, ExpMin(100.0, 100.8))
}

func TestExpMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Yusuf",
	}
	second := map[string]string{
		"Name": "Yusuf",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExpSlice(t *testing.T) {
	first := []string{"Manshur"}
	second := []string{"Manshur"}

	assert.True(t, slices.Equal(first, second))
}
