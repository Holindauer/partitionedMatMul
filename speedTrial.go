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

	fmt.Println("\nRunning speed trial on basic matrix multiplication algorithm")

	var times []time.Duration

	// iterate over matrix sizes
	for _, matrixSize := range matrixSizes {

		// create a square matrix
		var A *Matrix = NewMatrix(matrixSize, matrixSize, "randRange")
		var B *Matrix = NewMatrix(matrixSize, matrixSize, "randRange")

		// create output matrix
		var C *Matrix = NewMatrix(matrixSize, matrixSize, "zero")

		// start timer
		start := time.Now()

		// perform basic matrix multiplication
		BasicMatMul(A, B, C)

		// stop timer
		elapsed := time.Since(start)
		times = append(times, elapsed)

		// print time elapsed
		fmt.Printf("%dx%d: %s\n", matrixSize, matrixSize, elapsed)
	}

	return times
}

/**
 * @notice testBasicMatMulSpeed tests the speed of the basic matrix multiplication algorithm
 * on square matrices of varying sizes
 * @param t: the testing object
 */
func PartitionedMatMulSpeedTrial(matrixSizes []uint) []time.Duration {

	fmt.Println("\nRunning speed trial on partitioned matrix multiplication algorithm")

	var times []time.Duration

	// iterate over matrix sizes
	for _, matrixSize := range matrixSizes {

		// create a square matrix
		var A *Matrix = NewMatrix(matrixSize, matrixSize, "randRange")
		var B *Matrix = NewMatrix(matrixSize, matrixSize, "randRange")

		// create output matrix
		var C *Matrix = NewMatrix(matrixSize, matrixSize, "zero")

		// start timer
		start := time.Now()

		// perform basic matrix multiplication
		PartitionedMatMul(A, B, C)

		// stop timer
		elapsed := time.Since(start)
		times = append(times, elapsed)

		// print time elapsed
		fmt.Printf("%dx%d: %s\n", matrixSize, matrixSize, elapsed)
	}

	return times
}
