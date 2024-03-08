package main

/*
 * @notice this file contains an implementation of the basic matrix multiplication algorithm
 * @param A: the first matrix in the product
 * @param B: the second matrix in the product
 * @param C: the matrix to store the result of the product
 */
func BasicMatMul(A *Matrix, B *Matrix, C *Matrix) {

	// check if the matrices are compatible for multiplication
	if A.numCols != B.numRows {
		panic("Matrices are incompatible for multiplication")
	}

	// compute basic matrix multiplication algorithm
	var i, j, k uint

	// i'th row
	for i = 0; i < A.numRows; i++ {

		// j'th col
		for j = 0; j < B.numCols; j++ {

			// dot product i'th row w/ j'th col
			for k = 0; k < A.numCols; k++ {
				C.data[Index(i, j, C.numCols)] += A.data[Index(i, k, A.numCols)] * B.data[Index(k, j, B.numCols)]
			}
		}
	}
}
