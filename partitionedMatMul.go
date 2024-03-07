package main

/**
 * @notice this file contians an implementation of a partitioned matrix multiplication algorithm
 * where each partitioned matrix is computed concurrently. The goal is in increase the speed of
 * the basic matrix multiplication as much as possible.
 */

import (
	"sync"
)

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

/*
 * @notice matMul2x2() is a helper function for partitionedMatMul() that computes the matrix multiplication
 * of 2x2 partitions of the input matrices A and B. The result is stored in the output matrix C.
 * @dev The function accepts Matrix2x2 structs as input params, which are used to access elements of the input
 * matrices A and B. As well as modify the output matrix C.
 */
func matMul2x2(A Matrix2x2, B Matrix2x2, C Matrix2x2) {

	// compute the 2x2 matrix multiplication directly into C
	*C.a = (*A.a * *B.a) + (*A.b * *B.c)
	*C.b = (*A.a * *B.b) + (*A.b * *B.d)
	*C.c = (*A.c * *B.a) + (*A.d * *B.c)
	*C.d = (*A.c * *B.b) + (*A.d * *B.d)
}

/**
 * @notice partitionedMatMul() computes the matrix multiplication of two matrices A and B, and stores the result
 * in the output matrix C.
 * @dev The algorithm partitions the input matrices into 2x2 matrices, and computes the matrix multiplication of
 * each partition concurrently. The results are stored in the output matrix C directly duuring each goroutine.
 * @dev currently, this algorithm only supports square matrices of size 2^n x 2^n
 */
func partitionedMatMul(A *Matrix, B *Matrix, C *Matrix) {

	// check if the matrices are compatible for multiplication
	if A.cols != B.rows {
		panic("Matrices are incompatible for multiplication")
	}

	// determine the amount of partitions to split the matrix into
	// matrix will be split into as many 2x2 matrices as possible
	var rowPartitions uint = A.rows / 2
	var colPartitions uint = A.cols / 2

	// create wait group
	var wg sync.WaitGroup

	// iterate row partitions
	for i := 0; i < int(rowPartitions); i++ {

		// move row start
		rowStart := uint(i * 2)

		// iterate col partitions
		for j := 0; j < int(colPartitions); j++ {

			// move col start
			colStart := uint(j * 2)

			// increment wait group
			wg.Add(1)

			// allocate memory for matrix partitions
			var partitionA Matrix2x2 = Matrix2x2{
				&A.data[index(rowStart, colStart, A.cols)],
				&A.data[index(rowStart, colStart+1, A.cols)],
				&A.data[index(rowStart+1, colStart, A.cols)],
				&A.data[index(rowStart+1, colStart+1, A.cols)],
			}

			var partitionB Matrix2x2 = Matrix2x2{
				&B.data[index(colStart, rowStart, B.cols)],
				&B.data[index(colStart, rowStart+1, B.cols)],
				&B.data[index(colStart+1, rowStart, B.cols)],
				&B.data[index(colStart+1, rowStart+1, B.cols)],
			}

			var partitionC Matrix2x2 = Matrix2x2{
				&C.data[index(rowStart, colStart, C.cols)],
				&C.data[index(rowStart, colStart+1, C.cols)],
				&C.data[index(rowStart+1, colStart, C.cols)],
				&C.data[index(rowStart+1, colStart+1, C.cols)],
			}

			// launch goroutine for 2x2 MatMul
			go func(i int, j int) {

				// compute 2x2 matrix multiplication
				matMul2x2(partitionA, partitionB, partitionC)

				wg.Done() // decrement wait group
			}(i, j)
		}
	}

	wg.Wait() // wait for all goroutines to finish
}
