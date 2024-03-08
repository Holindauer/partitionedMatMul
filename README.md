# partitionedMatMul

# Problem Description

This repository contains an implementation of partitioned matrix multiplication for square matrices. The goal is to increase the efficiency of the basic sequential matrix multiplication algorithm:

$$
C_{ij} = \sum_{k=1}^{n} A_{ik} \cdot B_{kj}
$$

Where $A$, $B$ and $C$ are matrices of size $n \times n$.

This algorithm has a time complexity of $O(n^3)$

# Partitioned MatMul 

A matrix can be partitioned into submatrices, resulting in a block matrix. For example, if we have a a 4x4 matrix A:

$$
A = \begin{bmatrix}
a & b & c & d \\
e & f & g & h \\
i & j & k & l \\
m & n & o & p
\end{bmatrix}
$$

It can be partitioned into 4 submatrices, $A_{11}$, $A_{12}$, $A_{21}$ and $A_{22}$:

$$
A = \begin{bmatrix}
A_{11} & A_{12} \\
A_{21} & A_{22}
\end{bmatrix}
$$

Where:

$$
A_{11} = \begin{bmatrix} a & b \\ e & f \end{bmatrix}, \quad
A_{12} = \begin{bmatrix} c & d \\ g & h \end{bmatrix}, \quad
A_{21} = \begin{bmatrix} i & j \\ m & n \end{bmatrix}, \quad
A_{22} = \begin{bmatrix} k & l \\ o & p \end{bmatrix}
$$


Two partitioned matrices can then be computed in the same way as the original matrix, but with the submatrices as elements instead of scalars. 

$$
\begin{bmatrix}
A_{11} & A_{12} \\
A_{21} & A_{22}
\end{bmatrix}
\begin{bmatrix}
B_{11} & B_{12} \\
B_{21} & B_{22}
\end{bmatrix} =
\begin{bmatrix}
C_{11} & C_{12} \\
C_{21} & C_{22}
\end{bmatrix}
$$

The algorithm in this repository concurrently computes each partition in the resulting matrix $C$ 

# How to Run

This command will run the main function in main.go, which calls BasicMatmulSpeedTrial() and PartitionedMatmulSpeedTrial() from speedTrial.go. These functions will run matmuls on dense square matrices of size incrementing from 2 to 2048. The time taken for each matmul is recorded and printed to the console.


    go run .


Here are the results from running the program on my machine with dense matrice of random integer elements between 0 and 25:
    
### Basic MatMul

    2x2: 314ns
    4x4: 358ns
    8x8: 1.255µs
    16x16: 8.193µs
    32x32: 80.336µs
    64x64: 594.587µs
    128x128: 4.409978ms
    256x256: 34.159684ms
    512x512: 351.853245ms
    1024x1024: 2.953325329s
    2048x2048: 1m27.187725326s


### Partitioned MatMul

    2x2: 2.945µs
    4x4: 2.504µs
    8x8: 34.008µs
    16x16: 89.508µs
    32x32: 148.428µs
    64x64: 1.049667ms
    128x128: 4.847871ms
    256x256: 25.920338ms
    512x512: 112.281198ms
    1024x1024: 501.250832ms
    2048x2048: 9.366901927s


### execution time decrease:

    2x2: 2.631µs
    4x4: 2.146µs
    8x8: 32.753µs
    16x16: 81.315µs
    32x32: 68.092µs
    64x64: 455.08µs
    128x128: 437.893µs
    256x256: -8.239346ms
    512x512: -239.572047ms
    1024x1024: -2.452074497s
    2048x2048: -1m17.820823399s
