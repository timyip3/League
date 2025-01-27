package matrix

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func TestInvertMatrix(t *testing.T) {
	tests := []struct {
		name        string
		input       [][]string
		expected    [][]string
		expectError bool
	}{
		{
			name:        "Empty matrix",
			input:       [][]string{},
			expected:    nil,
			expectError: false,
		},
		{
			name: "1x1 matrix",
			input: [][]string{
				{"1"},
			},
			expected: [][]string{
				{"1"},
			},
			expectError: false,
		},
		{
			name: "2x2 matrix",
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			expected: [][]string{
				{"1", "3"},
				{"2", "4"},
			},
			expectError: false,
		},
		{
			name: "3x2 matrix",
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
				{"5", "6"},
			},
			expected: [][]string{
				{"1", "3", "5"},
				{"2", "4", "6"},
			},
			expectError: false,
		},
		{
			name: "Invalid matrix with inconsistent rows",
			input: [][]string{
				{"1", "2"},
				{"3"},
				{"5", "6"},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Matrix with empty row",
			input: [][]string{
				{},
				{},
			},
			expected:    [][]string{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := InvertMatrix(tt.input)

			// Check error cases
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check result
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InvertMatrix() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFlattenMatrix(t *testing.T) {
	tests := []struct {
		name        string
		input       [][]string
		expected    string
		expectError bool
	}{
		{
			name:        "Empty matrix",
			input:       [][]string{},
			expected:    "",
			expectError: false,
		},
		{
			name: "1x1 matrix",
			input: [][]string{
				{"1"},
			},
			expected:    "1\n",
			expectError: false,
		},
		{
			name: "2x2 matrix",
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			expected:    "1,2,3,4\n",
			expectError: false,
		},
		{
			name: "3x3 matrix",
			input: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			expected:    "1,2,3,4,5,6,7,8,9\n",
			expectError: false,
		},
		{
			name: "Matrix with inconsistent rows",
			input: [][]string{
				{"1", "2"},
				{"3"},
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Matrix with empty row",
			input: [][]string{
				{},
				{},
			},
			expected:    string("\n"),
			expectError: false,
		},
		{
			name: "Matrix with invalid characters",
			input: [][]string{
				{"1", "2,3"},
				{"4", "5"},
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Matrix with newline characters",
			input: [][]string{
				{"1", "2\n3"},
				{"4", "5"},
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Large matrix performance test",
			input: func() [][]string {
				matrix := make([][]string, 100)
				for i := range matrix {
					matrix[i] = make([]string, 100)
					for j := range matrix[i] {
						matrix[i][j] = strconv.Itoa(i*100 + j)
					}
				}
				return matrix
			}(),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FlattenMatrix(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				if tt.name != "Large matrix performance test" {
					t.Errorf("FlattenMatrix() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

// Benchmark for performance testing
func BenchmarkFlattenMatrix(b *testing.B) {
	matrix := make([][]string, 100)
	for i := range matrix {
		matrix[i] = make([]string, 100)
		for j := range matrix[i] {
			matrix[i][j] = strconv.Itoa(i*100 + j)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FlattenMatrix(matrix)
	}
}

func TestSumMatrix(t *testing.T) {
	tests := []struct {
		name        string
		matrix      [][]string
		expected    string
		expectError bool
	}{
		{
			name:        "Empty matrix",
			matrix:      [][]string{},
			expected:    "0",
			expectError: false,
		},
		{
			name: "Simple addition",
			matrix: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			expected:    "10",
			expectError: false,
		},
		{
			name: "Invalid number",
			matrix: [][]string{
				{"1", "abc"},
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Overflow test (addition)",
			matrix: [][]string{
				{strconv.FormatInt(math.MaxInt64, 10), "1"},
			},
			expected:    "9223372036854775808",
			expectError: false,
		},
		{
			name: "Negative numbers",
			matrix: [][]string{
				{"-1", "-2"},
				{"-3", "-4"},
			},
			expected:    "-10",
			expectError: false,
		},
		{
			name: "Large positive numbers",
			matrix: [][]string{
				{"9223372036854775808", "9223372036854775808"},
			},
			expected:    "18446744073709551616",
			expectError: false,
		},
		{
			name: "Large negative numbers",
			matrix: [][]string{
				{"-9223372036854775808", "-9223372036854775808"},
			},
			expected:    "-18446744073709551616",
			expectError: false,
		},
		{
			name: "Matrix with whitespace",
			matrix: [][]string{
				{" 1 ", " 2 "},
				{" 3 ", " 4 "},
			},
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SumMatrix(tt.matrix)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("SumMatrix() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark for performance testing
func BenchmarkAddMatrix(b *testing.B) {
	matrix := make([][]string, 100)
	for i := range matrix {
		matrix[i] = make([]string, 100)
		for j := range matrix[i] {
			matrix[i][j] = strconv.Itoa(i*100 + j)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SumMatrix(matrix)
	}
}

func TestMultiplyMatrix(t *testing.T) {
	tests := []struct {
		name        string
		matrix      [][]string
		expected    string
		expectError bool
	}{
		{
			name:        "Empty matrix",
			matrix:      [][]string{},
			expected:    "0",
			expectError: false,
		},
		{
			name: "Simple multiplication",
			matrix: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			expected:    "24",
			expectError: false,
		},
		{
			name: "Matrix with zero",
			matrix: [][]string{
				{"1", "0"},
				{"3", "4"},
			},
			expected:    "0",
			expectError: false,
		},
		{
			name: "Invalid number",
			matrix: [][]string{
				{"1", "abc"},
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Overflow test",
			matrix: [][]string{
				{strconv.FormatInt(math.MaxInt64, 10), "2"},
			},
			expected:    "18446744073709551614",
			expectError: false,
		},
		{
			name: "Mixed positive and negative",
			matrix: [][]string{
				{"-1", "2"},
				{"3", "-4"},
			},
			expected:    "24",
			expectError: false,
		},
		{
			name: "Large negative numbers",
			matrix: [][]string{
				{"-1000000", "-1000000"},
				{"1000000", "-1000000"},
			},
			expected:    "-1000000000000000000000000",
			expectError: false,
		},
		{
			name: "Negative one multiplication",
			matrix: [][]string{
				{"-1", "-1"},
				{"-1", "-1"},
			},
			expected:    "1",
			expectError: false,
		},
		{
			name: "Matrix with whitespace",
			matrix: [][]string{
				{" 1 ", " 2 "},
				{" 3 ", " 4 "},
			},
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MultiplyMatrix(tt.matrix)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("MultiplyMatrix() = %v, want %v", result, tt.expected)
			}
		})
	}
}
