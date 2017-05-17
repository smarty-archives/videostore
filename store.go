package videostore

import (
	"fmt"
	"strconv"
)

/**************************************************************************/

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

/**************************************************************************/

type Rental struct {
	Movie      *Movie
	DaysRented int
}

func (rental *Rental) determineAmount() (amount float64) {
	switch rental.Movie.PriceCode {
	case RegularMovie:
		amount += 2
		if rental.DaysRented > 2 {
			amount += (float64)(rental.DaysRented-2) * 1.5
		}
	case NewRelease:
		amount += (float64)(rental.DaysRented * 3)
	case ChildrensMovie:
		amount += 1.5
		if rental.DaysRented > 3 {
			amount += (float64)(rental.DaysRented-3) * 1.5
		}
	}
	return amount
}

func (rental *Rental) determinePoints() int {
	if rental.Movie.PriceCode == NewRelease && rental.DaysRented > 1 {
		return 2
	} else {
		return 1
	}
}

/**************************************************************************/

type RentalStatement struct {
	name   string
	total  float64
	points int
	body   string
}

func NewStatement(name string) *RentalStatement {
	return &RentalStatement{name: name}
}

func (this *RentalStatement) makeHeader() string {
	return "Rental Record for " + this.name + "\n"
}

func (this *RentalStatement) Include(rental *Rental) {
	amount := rental.determineAmount()
	this.total += amount
	this.points += rental.determinePoints()
	this.body += formatStatementLine(rental, amount)
}

func formatStatementLine(rental *Rental, amount float64) string {
	return "\t" + rental.Movie.Title + "\t" + fmt.Sprint(amount) + "\n"
}

func (this *RentalStatement) FormatStatement() string {
	return this.makeHeader() + this.body + this.makeFooter()
}

func (this *RentalStatement) makeFooter() string {
	return "You owed " + fmt.Sprint(this.total) + "\n" +
		"You earned " + strconv.Itoa(this.points) + " frequent renter points\n"
}

func (this *RentalStatement) AmountOwed() float64 {
	return this.total
}

func (this *RentalStatement) PointsEarned() int {
	return this.points
}
