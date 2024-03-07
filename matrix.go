package main

import (
	"math/rand"
)

// define matrix struct w/ 1D arr and dimensions
type Matrix struct {
	data []float64
	rows uint
	cols uint
}

/*
 * @notice The Matrix2x2 struct contains 4 pointers to the adresses of elements of a 2x2 matrix
 * @dev This struct is intended to be used within a goroutine for computing matmuls of 2x2
 * partitions of a larger matrix.
 */
type Matrix2x2 struct {
	a *float64
	b *float64
	c *float64
	d *float64
}

/**
 * @helper function to generate a random int between min and max
 */
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

/**
 * @notice this function initializes a matrix with random values between 0 and 25
 */
func MatrixInit(rows uint, cols uint, matType string) Matrix {

	// allocate memory for matrix
	var matrix Matrix = Matrix{
		make([]float64, rows*cols),
		rows,
		cols,
	}

	// init matrix data
	for i := 0; i < int(rows*cols); i++ {

		// initialize to the matType specificaiton
		if matType == "randRange" {
			matrix.data[i] = float64(RandomInt(randLowerBound, randUpperBound))

		} else if matType == "zero" {
			matrix.data[i] = 0
		}
	}

	return matrix
}

/**
 * @notice this helper functions computes the index stride for a 2D array stored in a 1D array
 */
func Index(i, j, cols uint) uint {
	return i*cols + j
}

/**
 * @notice matricesAreEqual() is a helper function that checks if two matrices are equal
 */
func MatricesAreEqual(A *Matrix, B *Matrix) bool {

	// check if the matrices have the same dimensions
	if A.rows != B.rows || A.cols != B.cols {
		return false
	}

	// iterate over the matrix
	for i := 0; i < int(A.rows); i++ {
		for j := 0; j < int(A.cols); j++ {

			if A.data[Index(uint(i), uint(j), A.cols)] != B.data[Index(uint(i), uint(j), B.cols)] {
				return false
			}
		}
	}

	return true
}
