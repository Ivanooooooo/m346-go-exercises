package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirtstName string
	LastName   string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	// TODO: embed full name and birth date information
	ProfileName      FullName
	BirthDateInfo    BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		ProfileName: FullName{
			FirtstName: "Ivano",
			LastName:   "Gavran Stojanovski",
		},
		BirthDateInfo: BirthDate{
			DayOfBirth:   1,
			MonthOfBirth: 12,
			YearOfBirth:  2007,
		},
		NumberOfSiblings: 1,   // TODO: adjust
		ZodiacSign:       'S', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
