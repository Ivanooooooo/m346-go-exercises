package main

func computeGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else if score >= 50 {
		return "E"
	} else {
		return "Keine F"
	}
}

func main() {
	grade := computeGrade(89)
	println(grade)
}
