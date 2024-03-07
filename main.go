package main

import (
	"fmt"
	"time"
)

// rand range for matrix element initialization
const randLowerBound int = 0
const randUpperBound int = 25

func main() {

	// square matrix sizes to test
	var matrixSizes []uint = []uint{2, 4, 8, 16, 32, 64, 128, 256, 512 /*, 1024, 2048, 4096, 8192*/}

	// run speed trial
	partitionedTimes := partitionedMatMulSpeedTrial(matrixSizes)
	basicTimes := basicMatMulSpeedTrial(matrixSizes)

	// compute time delta
	fmt.Println("Results:")
	var timeDelta []time.Duration = make([]time.Duration, len(matrixSizes))

	// iterate over matrix sizes
	for i := 0; i < len(matrixSizes); i++ {
		timeDelta[i] = partitionedTimes[i] - basicTimes[i]
		fmt.Printf("Matrix size: %dx%d, execution time decrease delta: %s\n", matrixSizes[i], matrixSizes[i], timeDelta[i])
	}
}
