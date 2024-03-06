package main

import (
	"fmt"
	"time"
)

// rand range for matrix element initialization
const randLowerBound int = 0
const randUpperBound int = 25

/**
 * @notice testBasicMatMulSpeed tests the speed of the basic matrix multiplication algorithm
 * on square matrices of varying sizes
 * @param t: the testing object
 */
func basicMatMulSpeedTrial() {

	fmt.Println("Running speed trial on basic matrix multiplication algorithm")

	// square matrix sizes to test
	var matrixSizes []uint = []uint{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

	// iterate over matrix sizes
	for _, matrixSize := range matrixSizes {

		// create a square matrix
		var A Matrix = matrixInit(matrixSize, matrixSize, "randRange")
		var B Matrix = matrixInit(matrixSize, matrixSize, "randRange")

		// create output matrix
		var C Matrix = matrixInit(matrixSize, matrixSize, "zero")

		// start timer
		start := time.Now()

		// perform basic matrix multiplication
		basicMatMul(&A, &B, &C)

		// stop timer
		elapsed := time.Since(start)

		// print time elapsed
		fmt.Printf("Basic matrix multiplication on %dx%d matrix took %s\n", matrixSize, matrixSize, elapsed)
	}
}
