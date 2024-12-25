package main

import (
	"fmt"
	"time"

	work "github.com/flccodes/goflc/work/readwrite"
)

func main() {
	start := time.Now() // Start of Processing time
	csvFile := "./work/example/salary.csv"
	out := work.FreadCSV(csvFile)
	fmt.Printf("Test: %v\n", out)

	// Processing time
	end := time.Now()

	elapsed := end.Sub(start) // time.Since(start) == time.Now.Sub(start)

	// end := time.Since(start)
	fmt.Printf("\nProcessing time: %v\n", elapsed)
}
