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
func matrixInit(rows uint, cols uint, matType string) Matrix {

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
			matrix.data[i] = float64(randomInt(randLowerBound, randUpperBound))

		} else if matType == "zero" {
			matrix.data[i] = 0
		}
	}

	return matrix
}

/**
 * @notice this helper functions computes the index stride for a 2D array stored in a 1D array
 */
func index(i, j, cols uint) uint {
	return i*cols + j
}
