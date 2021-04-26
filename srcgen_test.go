package srcgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsUpperCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "uppercase characters",
			input:    "ABC",
			expected: true,
		},
		{
			name:     "uppercase and lowercase characters",
			input:    "HiJ",
			expected: false,
		},
		{
			name:     "lowercase characters",
			input:    "xyz",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isUpperCase(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_SplitCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty",
			input:    "",
			expected: nil,
		},
		{
			name:     "single lowercase word",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "single uppercase word",
			input:    "WORLD",
			expected: []string{"WORLD"},
		},
		{
			name:     "two lower case words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "two initial upper case words",
			input:    "Hello World",
			expected: []string{"Hello", "World"},
		},
		{
			name:     "from snake_case",
			input:    "hello_world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "from kebab-case",
			input:    "hello-world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "from camelCase",
			input:    "helloWorld",
			expected: []string{"hello", "World"},
		},
		{
			name:     "from PascalCase",
			input:    "HelloWorld",
			expected: []string{"Hello", "World"},
		},
		{
			name:     "unknown case",
			input:    "src/${$parent}/test/${child}",
			expected: []string{"src/${$parent}/test/${child}"},
		},
		{
			name:     "empty part",
			input:    "A__B",
			expected: []string{"A", "B"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual := splitCase(tt.input)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		split    []string
		expected string
	}{
		{
			name:     "single lowercase word",
			split:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "single capitalized word",
			split:    []string{"Hello"},
			expected: "hello",
		},
		{
			name:     "single uppercase word",
			split:    []string{"HELLO"},
			expected: "hello",
		},
		{
			name:     "multiple lowercase words",
			split:    []string{"hello", "world"},
			expected: "hello_world",
		},
		{
			name:     "multiple capitalized words",
			split:    []string{"Hello", "World"},
			expected: "hello_world",
		},
		{
			name:     "multiple uppercase words",
			split:    []string{"HELLO", "WORLD"},
			expected: "hello_world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinToSnakeCase(tt.split)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ToKebabCase(t *testing.T) {
	tests := []struct {
		name     string
		split    []string
		expected string
	}{
		{
			name:     "single lowercase word",
			split:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "single capitalized word",
			split:    []string{"Hello"},
			expected: "hello",
		},
		{
			name:     "single uppercase word",
			split:    []string{"HELLO"},
			expected: "hello",
		},
		{
			name:     "multiple lowercase words",
			split:    []string{"hello", "world"},
			expected: "hello-world",
		},
		{
			name:     "multiple capitalized words",
			split:    []string{"Hello", "World"},
			expected: "hello-world",
		},
		{
			name:     "multiple uppercase words",
			split:    []string{"HELLO", "WORLD"},
			expected: "hello-world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinToKebabCase(tt.split)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ToCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		split    []string
		expected string
	}{
		{
			name:     "single lowercase word",
			split:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "single capitalized word",
			split:    []string{"Hello"},
			expected: "hello",
		},
		{
			name:     "single uppercase word",
			split:    []string{"HELLO"},
			expected: "hello",
		},
		{
			name:     "multiple lowercase words",
			split:    []string{"hello", "world"},
			expected: "helloWorld",
		},
		{
			name:     "multiple capitalized words",
			split:    []string{"Hello", "World"},
			expected: "helloWorld",
		},
		{
			name:     "multiple uppercase words",
			split:    []string{"HELLO", "WORLD"},
			expected: "helloWorld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinToCamelCase(tt.split)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ToGoCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		split    []string
		expected string
	}{
		{
			name:     "single lowercase word",
			split:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "single capitalized word",
			split:    []string{"Hello"},
			expected: "hello",
		},
		{
			name:     "single uppercase word",
			split:    []string{"HELLO"},
			expected: "hello",
		},
		{
			name:     "multiple lowercase words",
			split:    []string{"hello", "world"},
			expected: "helloWorld",
		},
		{
			name:     "multiple capitalized words",
			split:    []string{"Hello", "World"},
			expected: "helloWorld",
		},
		{
			name:     "multiple uppercase words",
			split:    []string{"HELLO", "WORLD"},
			expected: "helloWorld",
		},
		{
			name:     "id",
			split:    []string{"user", "id"},
			expected: "userID",
		},
		{
			name:     "url",
			split:    []string{"profile", "Url"},
			expected: "profileURL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinToGoCamelCase(tt.split)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ToPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		split    []string
		expected string
	}{
		{
			name:     "single lowercase word",
			split:    []string{"hello"},
			expected: "Hello",
		},
		{
			name:     "single capitalized word",
			split:    []string{"Hello"},
			expected: "Hello",
		},
		{
			name:     "single uppercase word",
			split:    []string{"HELLO"},
			expected: "Hello",
		},
		{
			name:     "multiple lowercase words",
			split:    []string{"hello", "world"},
			expected: "HelloWorld",
		},
		{
			name:     "multiple capitalized words",
			split:    []string{"Hello", "World"},
			expected: "HelloWorld",
		},
		{
			name:     "multiple uppercase words",
			split:    []string{"HELLO", "WORLD"},
			expected: "HelloWorld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinToPascalCase(tt.split)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_ToGoPascalCase(t *testing.T) {
	tests := []struct {
		name     string
		split    []string
		expected string
	}{
		{
			name:     "single lowercase word",
			split:    []string{"hello"},
			expected: "Hello",
		},
		{
			name:     "single capitalized word",
			split:    []string{"Hello"},
			expected: "Hello",
		},
		{
			name:     "single uppercase word",
			split:    []string{"HELLO"},
			expected: "Hello",
		},
		{
			name:     "multiple lowercase words",
			split:    []string{"hello", "world"},
			expected: "HelloWorld",
		},
		{
			name:     "multiple capitalized words",
			split:    []string{"Hello", "World"},
			expected: "HelloWorld",
		},
		{
			name:     "multiple uppercase words",
			split:    []string{"HELLO", "WORLD"},
			expected: "HelloWorld",
		},
		{
			name:     "id",
			split:    []string{"id"},
			expected: "ID",
		},
		{
			name:     "url",
			split:    []string{"Url"},
			expected: "URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := joinToGoPascalCase(tt.split)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
