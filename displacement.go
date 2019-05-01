package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.0
	vi := 0.0
	si := 0.0

	fmt.Println("Please enter acceleration")
	fmt.Scan(&a)
	fmt.Println("Please enter initial velocity")
	fmt.Scan(&vi)
	fmt.Println("Please enter initial displacement")
	fmt.Scan(&si)

	s := GenDisplaceFn(a, vi, si)

	fmt.Println(s(a, vi, si))
}

func GenDisplaceFn(a, vi, si float64) func(float64, float64, float64) float64 {
	fn := func(a, vi, si float64) (t float64) {
		fmt.Println("Please enter time")
		fmt.Scan(&t)
		t = (1/2*a)*math.Pow(t, 2) + (vi * t) + si
		return t
	}
	return fn
}
