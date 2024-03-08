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
	var matrixSizes []uint = []uint{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	// run speed trial
	partitionedTimes := PartitionedMatMulSpeedTrial(matrixSizes)
	basicTimes := BasicMatMulSpeedTrial(matrixSizes)

	// compute time delta
	fmt.Println("\nResults (execution time decrease):")
	var timeDelta []time.Duration = make([]time.Duration, len(matrixSizes))

	// iterate over matrix sizes
	for i := 0; i < len(matrixSizes); i++ {
		timeDelta[i] = partitionedTimes[i] - basicTimes[i]
		fmt.Printf("%dx%d: %s\n", matrixSizes[i], matrixSizes[i], timeDelta[i])
	}
}
