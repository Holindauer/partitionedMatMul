package main

/**
 * @notice this file contians an implementation of a partitioned matrix multiplication algorithm
 * where each partitioned matrix is computed concurrently. The goal is in increase the speed of
 * the basic matrix multiplication as much as possible.
 */

import (
	"fmt"
	"sync"
)

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
 *	@notice
 */

/**
 * @notice partitionedMatMul() computes the matrix multiplication of two matrices A and B, and stores the result
 * in the output matrix C.
 * @dev The algorithm partitions the input matrices into 2x2 matrices, and computes the matrix multiplication of
 * each partition concurrently. The results are stored in the output matrix C directly duuring each goroutine.
 * @dev currently, this algorithm only supports square matrices of size 2^n x 2^n
 */
func PartitionedMatMul(A *Matrix, B *Matrix, C *Matrix) {

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

	// ! new Game Plan
	/**
	* After writing the test in MatMul_test.go, I realized that this function is not actually computing a partitioned
	* matmul. Currentlly, it is apply elementwise multiplications on the partitions of the input matrices, and storing
	* the results in the output matrix. This is not the correct algorithm.
	*
	* The algorithm needs to be refactored to:
	*
	* 1.) determine partitions (for the most part, already done but may need to be adjusted)
	* 2.) iterate over partitions of the output matrix (^^)
	* 3.) compute the elements of output concurrentlly by computing each matrix partition in the output matrix. Since each
	*     of the partitions in the output represents a seperate dot product specific to certain rows and columns of the two
	*     input matrices, these can be computed concurrently in a two step process...
	*
	*     Two Step Process:
	*     a.) A groutine is launched to compute the output matrix partition for each partition within the output. Each
	          goroutine, pass a pointer to the original matrix that will NOT be modified, it is being passed in to
	*         access the elements of the input matrices. And to assemble further partitions during computation of the output

	*     b.) Inside the above described goroutine, for the specific partition of the output, the indices of each of the
	          individual 2x2 matrix partitions involved in the dot product computation (involving the input and output mat)
			  will be computed, and a goroutine will be launched to compute the matmul of each partition multiplication. Their
			  result will be elementwise added into the output matrix using a MuTex to lock shared resources and prevent race
			  conditions.
	*

	*/

	// ! below to be refactored

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
				&A.data[Index(rowStart, colStart, A.cols)],
				&A.data[Index(rowStart, colStart+1, A.cols)],
				&A.data[Index(rowStart+1, colStart, A.cols)],
				&A.data[Index(rowStart+1, colStart+1, A.cols)],
			}

			var partitionB Matrix2x2 = Matrix2x2{
				&B.data[Index(colStart, rowStart, B.cols)],
				&B.data[Index(colStart, rowStart+1, B.cols)],
				&B.data[Index(colStart+1, rowStart, B.cols)],
				&B.data[Index(colStart+1, rowStart+1, B.cols)],
			}

			var partitionC Matrix2x2 = Matrix2x2{
				&C.data[Index(rowStart, colStart, C.cols)],
				&C.data[Index(rowStart, colStart+1, C.cols)],
				&C.data[Index(rowStart+1, colStart, C.cols)],
				&C.data[Index(rowStart+1, colStart+1, C.cols)],
			}

			// launch goroutine for 2x2 MatMul
			go func(i int, j int) {

				// compute 2x2 matrix multiplication
				matMul2x2(partitionA, partitionB, partitionC)

				fmt.Println(" C output matrix: ", partitionC.a, partitionC.b, partitionC.c, partitionC.d)

				wg.Done() // decrement wait group
			}(i, j)
		}
	}

	wg.Wait() // wait for all goroutines to finish
}
