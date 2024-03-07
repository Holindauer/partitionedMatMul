package main

import (
	"fmt"
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

		fmt.Printf("Testing matrix %dx%d matrix\n", matrixSize, matrixSize)

		// create a square matrix
		var A Matrix = MatrixInit(matrixSize, matrixSize, "randRange")
		var B Matrix = MatrixInit(matrixSize, matrixSize, "randRange")

		fmt.Println(" A input matrix: ", A.data)
		fmt.Println(" B input matrix: ", B.data)

		// create output matrices
		var C_basic Matrix = MatrixInit(matrixSize, matrixSize, "zero")
		var C_partitioned Matrix = MatrixInit(matrixSize, matrixSize, "zero")

		fmt.Println(" C basic output matrix: ", C_basic.data)
		fmt.Println(" C partitioned output matrix: ", C_partitioned.data)

		// perform basic matrix multiplication
		BasicMatMul(&A, &B, &C_basic)

		fmt.Println(" C basic output matrix: ", C_basic.data)
		fmt.Println(" C partitioned output matrix: ", C_partitioned.data)

		// perform partitioned matrix multiplication
		PartitionedMatMul(&A, &B, &C_partitioned)

		fmt.Println(" C partitioned output matrix: ", C_partitioned.data)

		// check if the results are the same
		if !MatricesAreEqual(&C_basic, &C_partitioned) {
			panic("Matrices are not equal")
		}
	}
}
