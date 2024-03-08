package main

import (
	"math/rand"
	"sync"
)

// Matrix structure to hold the matrix data
type Matrix struct {
	data    []float64
	numRows uint
	numCols uint
	mu      sync.Mutex // A single mutex for the whole matrix
}

// Matrix2x2 structure for partition calculation
type Matrix2x2 struct {
	a, b, c, d *float64
}

// Initialize a new matrix with given dimensions and type
func NewMatrix(rows, cols uint, matType string) *Matrix {
	matrix := &Matrix{
		data:    make([]float64, rows*cols),
		numRows: rows,
		numCols: cols,
	}

	for i := range matrix.data {
		if matType == "randRange" {
			matrix.data[i] = float64(rand.Intn(10)) // Simplified for brevity
		} else if matType == "zero" {
			matrix.data[i] = 0
		} else {
			panic("Invalid matType specified")
		}
	}

	return matrix
}

// @helper this helper functions computes the index stride for a 2D array stored in a 1D array
func Index(i, j, cols uint) uint {
	return i*cols + j
}

// @helper matricesAreEqual() is a helper function that checks if two matrices are equal
func MatricesAreEqual(A *Matrix, B *Matrix) bool {

	// check if the matrices have the same dimensions
	if A.numRows != B.numRows || A.numCols != B.numCols {
		return false
	}

	// iterate over the matrix
	for i := 0; i < int(A.numRows); i++ {
		for j := 0; j < int(A.numCols); j++ {

			// check if the elements are equal
			if A.data[Index(uint(i), uint(j), A.numCols)] != B.data[Index(uint(i), uint(j), B.numCols)] {
				return false
			}
		}
	}

	return true
}

// @helper function to generate a random int between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
