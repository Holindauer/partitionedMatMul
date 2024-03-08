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
\end{bmatrix}
=
\begin{bmatrix}
C_{11} & C_{12} \\
C_{21} & C_{22}
\end{bmatrix}
$$

The algorithm in this repository concurrently computes each partition in the resulting matrix $C$ 