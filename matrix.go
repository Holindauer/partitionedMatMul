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

/**
 * @helper function to generate a random int between min and max
 */
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

/**
 * @notice this function initializes a matrix with random values between 0 and 25
 */
func matrixInit(rows uint, cols uint) Matrix {
	// init matrix struct
	var matrix Matrix = Matrix{
		make([]float64, rows*cols),
		rows,
		cols,
	}

	for i := 0; i < int(rows*cols); i++ {
		// random int between 0 and 25
		matrix.data[i] = float64(randomInt(0, 25))
	}

	return matrix
}

/**
 * @notice this helper functions computes the index stride for a 2D array stored in a 1D array
 */
func index(i, j, cols uint) uint {
	return i*cols + j
}
