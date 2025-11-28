package main

import "fmt"

func main() {
	type Student struct {
		FirstName string
		LastName  string
	}

	type Class struct {
		Students map[byte]Student
	}

	modules := map[string]Class{
		"M117": {
			Students: map[byte]Student{
				1: {FirstName: "Ivano", LastName: "Gavran Stojanovski"},
				2: {FirstName: "Nino", LastName: "Meier"},
			},
		},
		"M346": {
			Students: map[byte]Student{
				1: {FirstName: "Cristiano", LastName: "Ronaldo"},
				2: {FirstName: "Lionel", LastName: "Messi"},
			},
		},
		"M104": {
			Students: map[byte]Student{
				1: {FirstName: "Lewis", LastName: "Hamilton"},
				2: {FirstName: "Sebastian", LastName: "Vettel"},
			},
		},
	}

	fmt.Println(modules)
}
