// create main
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")

	var A Matrix = matrixInit(1000, 1000)
	var B Matrix = matrixInit(1000, 1000)
	var C Matrix = matrixInit(1000, 1000)

	basicMatMul(&A, &B, &C)

	// fmt.Println(C)
}
