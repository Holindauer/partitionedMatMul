package main

import (
	"testing"
)

/**
 * @test testSameMatMulResult tests that the basic matrix multiplication algorithm
 * and the partitioned matrix multiplication algorithm produce the same result
 * @param t: the testing object
 */
func Test_SameMatMulResult(t *testing.T) {

	// square matrix sizes to test
	var matrixSizes []uint = []uint{2, 4, 8, 16, 32, 64, 128, 256, 512}

	// iterate over matrix sizes
	for _, matrixSize := range matrixSizes {

		// create a square matrix
		var A *Matrix = NewMatrix(matrixSize, matrixSize, "randRange")
		var B *Matrix = NewMatrix(matrixSize, matrixSize, "randRange")

		// create output matrices
		var C_basic *Matrix = NewMatrix(matrixSize, matrixSize, "zero")
		var C_partitioned *Matrix = NewMatrix(matrixSize, matrixSize, "zero")

		// perform basic matrix multiplication
		BasicMatMul(A, B, C_basic)

		// perform partitioned matrix multiplication
		PartitionedMatMul(A, B, C_partitioned)

		// check if the results are the same
		if !MatricesAreEqual(C_basic, C_partitioned) {
			panic("Matrices are not equal")
		}
	}
}
