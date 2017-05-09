package videostore

import (
	"fmt"
	"strconv"
)

type PriceCode int

const (
	RegularMovie PriceCode = iota
	ChildrensMovie
	NewRelease
)

type Movie struct {
	Title     string
	PriceCode PriceCode
}

type Rental struct {
	Movie      *Movie
	DaysRented int
}

type Customer struct {
	name    string
	rentals []*Rental
}

func (this *Customer) Add(rental *Rental) {
	this.rentals = append(this.rentals, rental)
}

func (this *Customer) Statement() string {
	total := 0.0
	points := 0
	result := "Rental Record for " + this.name + "\n"

	for i := 0; i < len(this.rentals); i++ {
		amt := 0.0
		r := this.rentals[i]

		// determine the amount for each line
		switch r.Movie.PriceCode {
		case RegularMovie:
			amt += 2
			if r.DaysRented > 2 {
				amt += (float64)(r.DaysRented-2) * 1.5
			}
		case NewRelease:
			amt += (float64)(r.DaysRented * 3)
		case ChildrensMovie:
			amt += 1.5
			if r.DaysRented > 3 {
				amt += (float64)(r.DaysRented-3) * 1.5
			}
		}

		points++

		if r.Movie.PriceCode == NewRelease && r.DaysRented > 1 {
			points++
		}

		result += "\t" + r.Movie.Title + "\t" + fmt.Sprint(amt) + "\n"
		total += amt
	}

	result += "You owed " + fmt.Sprint(total) + "\n"
	result += "You earned " + strconv.Itoa(points) + " frequent renter points\n"

	return result
}
