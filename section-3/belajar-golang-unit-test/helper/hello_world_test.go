package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarkCases := []struct {
		name    string
		request string
	}{
		{
			name:    "Yusuf",
			request: "Yusuf",
		},
		{
			name:    "MuhammadYusufManshur",
			request: "Muhammad Yusuf Manshur",
		},
	}
	for _, bC := range benchmarkCases {
		b.Run(bC.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bC.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Yusuf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Yusuf")
		}
	})
	b.Run("Manshur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Manshur")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Yusuf")
	}
}

func BenchmarkHelloWorldManshur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Manshur")
	}
}

func TestTable(t *testing.T) {
	testCases := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Yusuf",
			request:  "Yusuf",
			expected: "Hello Yusuf",
		},
		{
			name:     "Manshur",
			request:  "Manshur",
			expected: "Hello Manshur",
		},
		{
			name:     "Muhammad",
			request:  "Muhammad",
			expected: "Hello Muhammad",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := HelloWorld(tC.request)
			require.Equal(t, tC.expected, result)
		})
	}
}

func TestSub(t *testing.T) {
	t.Run("Yusuf", func(t *testing.T) {
		result := HelloWorld("Yusuf")
		require.Equal(t, "Hello Yusuf", result, "Result must be 'Hello Yusuf'")
	})
	t.Run("Manshur", func(t *testing.T) {
		result := HelloWorld("Manshur")
		require.Equal(t, "Hello Manshur", result, "Result must be 'Hello Manshur'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

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
