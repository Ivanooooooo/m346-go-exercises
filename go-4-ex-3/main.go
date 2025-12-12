package main

import "math"

func computeQuadraticFormula(a float64, b float64, c float64) (float64, float64) {
	positive := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	negative := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	return positive, negative
}

func main() {
	positiveLösung, negativeLösung := computeQuadraticFormula(1, -3, 2)
	println(positiveLösung, "+-", negativeLösung)
}
