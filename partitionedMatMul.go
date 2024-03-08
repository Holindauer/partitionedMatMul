package main

/**
 * @notice partitionedMatMul.go contians an implementation of partitioned matrix multiplication for two
 * Matrix structs. Two matrices A and B are partitioned into 2x2 matrices, and their matrix multiplication
 * is computed concurrently, and stored within an Output Matrix.
 * @goal The goal of this algorithm is to reduce time complexity of the basic matmulmul algorithm as much
 * as possible using the partitioned matmul approach.
 */

import (
	"sync"
)

/**
 * @notice PartitionedMatMul() computes the product of two matrices A and B using the partitioned matrix
 * multiplication algorithm. The result is stored in the Output matrix.
 * @param A: the first matrix in the product
 * @param B: the second matrix in the product
 * @param Output: the matrix to store the result of the product
 */
func PartitionedMatMul(A, B, Output *Matrix) {
	if A.numCols != B.numRows {
		panic("Matrices are incompatible for multiplication")
	}

	partitions := A.numRows / 2
	var wg sync.WaitGroup

	// Compute each output partition concurrently
	for i := uint(0); i < partitions; i++ {
		for j := uint(0); j < partitions; j++ {

			wg.Add(1)
			go func(i, j uint) {
				ComputeOutputPartition(A, B, Output, i*2, j*2)
				wg.Done()
			}(i, j)
		}
	}

	wg.Wait() // Wait for all goroutines to finish
}

/*
 * @notice ComputeOutputPartition() computes a single 2x2 partition of the Output matrix for the specified
 * row and column start indices. This function is called concurrently by the PartitionedMatMul() function.
 * @param A: the first matrix in the product
 * @param B: the second matrix in the product
 * @param Output: the matrix to store the result of the product
 * @param rowStart: the starting row index of the partition
 * @param colStart: the starting column index of the partition
 */
func ComputeOutputPartition(A, B, Output *Matrix, rowStart, colStart uint) {

	subMatSize := uint(2) // Assuming 2x2 submatrices

	// Iterate over "block matrices"
	for i := rowStart; i < rowStart+subMatSize; i++ {
		for j := colStart; j < colStart+subMatSize; j++ {

			// Compute each element in the 2x2 partition
			sum := float64(0)
			for k := uint(0); k < A.numCols; k++ {
				sum += A.data[i*A.numCols+k] * B.data[k*B.numCols+j]
			}

			// Safely update the output matrix
			Output.mu.Lock()
			Output.data[i*Output.numCols+j] = sum
			Output.mu.Unlock()
		}
	}
}
