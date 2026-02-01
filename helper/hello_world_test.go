package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yudha")

	if result != "Hello Yudha" {
		panic("Result is not 'Hello Yudha'")
	}
}

func TestHelloWorldDede(t *testing.T) {
	result := HelloWorld("Dede")

	if result != "Hello Dede" {
		panic("Result is not 'Hello Dede'")
	}
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Yudha")

	if result != "Hello Yudha Fail" {
		t.Fail()
	}

	fmt.Println("Keep running Fail")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Yudha")

	if result != "Hello World Yudha Fail Now" {
		t.FailNow()
	}

	fmt.Println("Keep Runnig Fail Now")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Yudha")

	if result != "Hello World Yudha Error" {
		t.Error("Error message harusnya xxxx")
	}

	fmt.Println("Keep Running Error")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Yudha")

	if result != "Hello World Yudha Fatal" {
		t.Fatal("Error message harusnya xxxx")
	}

	fmt.Println("Keep Running Fatal")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Yudha")
	assert.Equal(t, "Hello YudhaL", result, "Hasilnya harus xxxx")

	fmt.Println("Keep running assertion")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Yudha")
	require.Equal(t, "Hello YudhaL", result, "Hasilnya harus xxxx")

	fmt.Println("Keep running require")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Only run on mac OS")
	}

	fmt.Println("Test logic running")
}

func TestSubTest(t *testing.T) {
	t.Run("Test Yudha", func(t *testing.T) {
		result := HelloWorld("Yudha")
		require.Equal(t, "Hello Yudha", result, "Result must be 'Hello Yudha'")
	})

	t.Run("Test Dede", func(t *testing.T) {
		result := HelloWorld("Dede")
		require.Equal(t, "Hello Laksmana", result, "Result not satisfied")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name: "Test Yudha",
			request: "Yudha",
			expected: "Hello Yudha",
		},
		{
			name: "Test Laksmana",
			request: "Laksmana",
			expected: "Hello Laksmana",
		},
		{
			name: "Test Dede",
			request: "Dede",
			expected: "Hello Dede",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("before test")

	m.Run()

	fmt.Println("after test")
}