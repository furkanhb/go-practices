package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"fist":   13,
		"second": 45,
	}

	floats := map[string]float64{
		"fist":   42.32,
		"second": 45.55,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, num := range m {
		s += num
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, num := range m {
		s += num
	}

	return s
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
