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
	total float64
	points int
}

func (this *Customer) Add(rental *Rental) {
	this.rentals = append(this.rentals, rental)
}

func (this *Customer) Statement() string {
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

		this.points++

		if r.Movie.PriceCode == NewRelease && r.DaysRented > 1 {
			this.points++
		}

		result += "\t" + r.Movie.Title + "\t" + fmt.Sprint(amt) + "\n"
		this.total += amt
	}

	result += "You owed " + fmt.Sprint(this.total) + "\n"
	result += "You earned " + strconv.Itoa(this.points) + " frequent renter points\n"

	return result
}

func (this *Customer) AmountOwed() float64 {
	return this.total
}

func (this *Customer) PointsEarned() int {
	return this.points
}
