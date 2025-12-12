package main

import "math"

func computeHypotenuse(a float64, b float64) float64 {
	ab := math.Pow(a, 2) + math.Pow(b, 2)
	return math.Sqrt(ab)

}

func main() {
	hypotenuse := computeHypotenuse(3, 4)
	print(hypotenuse)

}
