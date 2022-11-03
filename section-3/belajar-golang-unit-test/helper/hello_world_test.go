package helper

import (
	"fmt"
	"testing"
)

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
