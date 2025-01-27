package matrix

import (
	"fmt"
	"math/big"
	"strings"
)

func InvertMatrix(matrix [][]string) ([][]string, error) {
	if len(matrix) == 0 {
		return nil, nil
	}

	rows := len(matrix)
	cols := len(matrix[0])

	if rows == 1 && cols == 1 {
		return [][]string{{matrix[0][0]}}, nil
	}

	inverted := make([][]string, cols)
	for i := range inverted {
		inverted[i] = make([]string, rows)
	}

	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			if len(matrix[i]) != cols {
				return nil, fmt.Errorf("invalid matrix: inconsistent rows and columns")
			}
			// Validate number
			_, ok := new(big.Int).SetString(matrix[i][j], 10)
			if !ok {
				return nil, fmt.Errorf("invalid number at position [%d,%d]", i, j)
			}

			inverted[j][i] = matrix[i][j]
		}
	}

	return inverted, nil
}

func FlattenMatrix(matrix [][]string) (string, error) {
	if len(matrix) == 0 {
		return "", nil
	}

	rows := len(matrix)
	if rows == 0 {
		return "", nil
	}

	cols := len(matrix[0])

	if rows == 1 && cols == 1 {
		return matrix[0][0] + "\n", nil
	}

	// Calculate exact capacity needed
	// Formula: sum of lengths of all strings + (rows*cols - 1) commas + 1 newline
	totalCap := 0
	for _, row := range matrix {
		if len(row) != cols {
			return "", fmt.Errorf("invalid matrix: inconsistent rows and columns")
		}
		for _, val := range row {
			totalCap += len(val)
		}
	}
	totalCap += (rows*cols - 1) + 1 // Add space for commas and newline

	// Initialize builder with exact capacity
	var flattenBuilder strings.Builder
	flattenBuilder.Grow(totalCap)

	// Flatten matrix efficiently
	for i, row := range matrix {
		for j, val := range row {
			// Validate string content
			if strings.ContainsAny(val, ",\n") {
				return "", fmt.Errorf("invalid character in matrix at position [%d,%d]: value contains comma or newline", i, j)
			}

			// Validate number
			_, ok := new(big.Int).SetString(matrix[i][j], 10)
			if !ok {
				return "", fmt.Errorf("invalid number at position [%d,%d]", i, j)
			}

			flattenBuilder.WriteString(val)

			// Add comma if not the last element
			if !(i == rows-1 && j == cols-1) {
				flattenBuilder.WriteByte(',')
			}
		}
	}

	flattenBuilder.WriteByte('\n')
	return flattenBuilder.String(), nil
}

func SumMatrix(matrix [][]string) (string, error) {
	if len(matrix) == 0 {
		return "0", nil
	}

	cols := len(matrix[0])
	if cols == 0 {
		return "", fmt.Errorf("invalid matrix: empty row found")
	}

	// Initialize result based on operation
	result := big.NewInt(0)

	// Process matrix elements
	for i, row := range matrix {
		for j, val := range row {
			if len(row) != cols {
				return "", fmt.Errorf("invalid matrix: inconsistent row length at row %d", i)
			}

			// Validate and parse number
			integer, ok := new(big.Int).SetString(val, 10)
			if !ok {
				return "", fmt.Errorf("invalid number at position [%d,%d]", i, j)
			}

			result.Add(result, integer)
		}
	}

	return result.Text(10), nil
}

func MultiplyMatrix(matrix [][]string) (string, error) {
	if len(matrix) == 0 {
		return "0", nil
	}

	cols := len(matrix[0])
	if cols == 0 {
		return "", fmt.Errorf("invalid matrix: empty row found")
	}

	// Initialize result
	result := big.NewInt(1)

	// Process matrix elements
	for i, row := range matrix {
		for j, val := range row {
			if len(row) != cols {
				return "", fmt.Errorf("invalid matrix: inconsistent row length at row %d", i)
			}

			// Validate and parse number
			integer, ok := new(big.Int).SetString(val, 10)
			if !ok {
				return "", fmt.Errorf("invalid number at position [%d,%d]", i, j)
			}

			// 	If multiplying by zero, return early
			if integer.Cmp(big.NewInt(0)) == 0 {
				return "0", nil
			}

			result.Mul(result, integer)
		}
	}

	return result.Text(10), nil
}
