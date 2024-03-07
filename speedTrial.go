package main

import (
	"fmt"
	"time"
)

/**
 * @notice testBasicMatMulSpeed tests the speed of the basic matrix multiplication algorithm
 * on square matrices of varying sizes
 * @param t: the testing object
 */
func BasicMatMulSpeedTrial(matrixSizes []uint) []time.Duration {

	fmt.Println("Running speed trial on basic matrix multiplication algorithm")

	var times []time.Duration

	// iterate over matrix sizes
	for _, matrixSize := range matrixSizes {

		// create a square matrix
		var A Matrix = MatrixInit(matrixSize, matrixSize, "randRange")
		var B Matrix = MatrixInit(matrixSize, matrixSize, "randRange")

		// create output matrix
		var C Matrix = MatrixInit(matrixSize, matrixSize, "zero")

		// start timer
		start := time.Now()

		// perform basic matrix multiplication
		BasicMatMul(&A, &B, &C)

		// stop timer
		elapsed := time.Since(start)
		times = append(times, elapsed)

		// print time elapsed
		fmt.Printf("Basic matrix multiplication on %dx%d matrix took %s\n", matrixSize, matrixSize, elapsed)
	}

	return times
}

/**
 * @notice testBasicMatMulSpeed tests the speed of the basic matrix multiplication algorithm
 * on square matrices of varying sizes
 * @param t: the testing object
 */
func PartitionedMatMulSpeedTrial(matrixSizes []uint) []time.Duration {

	fmt.Println("Running speed trial on basic matrix multiplication algorithm")

	var times []time.Duration

	// iterate over matrix sizes
	for _, matrixSize := range matrixSizes {

		// create a square matrix
		var A Matrix = MatrixInit(matrixSize, matrixSize, "randRange")
		var B Matrix = MatrixInit(matrixSize, matrixSize, "randRange")

		// create output matrix
		var C Matrix = MatrixInit(matrixSize, matrixSize, "zero")

		// start timer
		start := time.Now()

		// perform basic matrix multiplication
		PartitionedMatMul(&A, &B, &C)

		// stop timer
		elapsed := time.Since(start)
		times = append(times, elapsed)

		// print time elapsed
		fmt.Printf("Basic matrix multiplication on %dx%d matrix took %s\n", matrixSize, matrixSize, elapsed)
	}

	return times
}
