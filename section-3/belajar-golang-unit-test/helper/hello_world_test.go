package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Yusuf")
	require.Equal(t, "Hello Yusuf", result, "Result must be 'Hello Yusuf'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Yusuf")
	require.Equal(t, "Hello Yusuf", result, "Result must be 'Hello Yusuf'")
	fmt.Println("TestHelloWorldRequire Done")

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Yusuf")
	assert.Equal(t, "Hello Yusuf", result, "Result must be 'Hello Yusuf'")
	fmt.Println("TestHelloWorldAssert Done")

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yusuf")
	if result != "Hello Yusuf" {
		// unit test failed
		// panic("Result is not 'Hello Yusuf'")
		// t.Fail()
		t.Error("Result must be 'Hello Yusuf'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldManshur(t *testing.T) {
	result := HelloWorld("Manshur")
	if result != "Hello Manshur" {
		// unit test failed
		// panic("Result is not 'Hello Manshur'")
		// t.FailNow()
		t.Fatal("Result must be 'Hello Manshur'")
	}

	fmt.Println("TestHelloWorldManshur Done")
}
