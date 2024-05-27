package main

import (
	"fmt"

	"github.com/TheFoxForever/assignment_week_9/trimmedmean"
)

func main() {
	// proportion to be trimmed from both ends
	proportion := 0.05

    // lists of data that hold integer and float values, at 100 each
	intData := make([]float64, 100)
    for i := 0; i < 100; i++ {
        intData[i] = float64(i + 1)
    }

    floatData := make([]float64, 100)
	for i := 0; i < 100; i++ {
		floatData = append(floatData, float64(i+1)+0.5)
	}

	// calculations
    intMean, err := trimmedmean.TrimmedMean(intData, proportion)
	//if error, then we show the error
    if err != nil {
        fmt.Println("Error calculating trimmed mean:", err)
    } else {
        fmt.Println("Trimmed Mean:", intMean)
    }

	//and for float values
    floatMean, err := trimmedmean.TrimmedMean(floatData, proportion)
    if err != nil {
        fmt.Println("Error calculating trimmed mean:", err)
    } else {
        fmt.Println("Trimmed Mean:", floatMean)
    }
}
