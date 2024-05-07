package tutgenerics

/*
Go Tutorial: Getting started with generics
Source: https://go.dev/doc/tutorial/generics
*/

import (
	"fmt"
)

var ints = map[string]int64{
	"first":  4,
	"second": 2,
}

var floats = map[string]float64{
	"first":  8.5,
	"second": 24.5,
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	// Go Tutorials: Generics: https://go.dev/doc/tutorial/generics
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

// FcallnonGen initializes the Non-Generics Functions
func FcallNonGen() {

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

}

// SumIntsOrSumFloats sums the values of the map m. It supports both int64 and float64
// as type for map values
func SumIntsOrSumFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// FcallGen initializes the generic function
func FcallGen() {
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrSumFloats[string, int64](ints),
		SumIntsOrSumFloats[string, float64](floats))
}
