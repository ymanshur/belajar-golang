package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yusuf")
	if result != "Hello Yusuf" {
		// unit test failed
		panic("Result is not 'Hello Yusuf'")
	}
}

func TestHelloWorldManshur(t *testing.T) {
	result := HelloWorld("Manshur")
	if result != "Hello Manshur" {
		// unit test failed
		panic("Result is not 'Hello Manshur'")
	}
}
