package videostore

import "fmt"

type movie interface {
	determineAmount(daysRented int) float64
	determinePoints(daysRented int) int
}

type Rental struct {
	movie      movie
	daysRented int
}

func NewRental(movie movie, daysRented int) *Rental {
	return &Rental{movie: movie, daysRented: daysRented}
}

func (rental Rental) DetermineAmount() (amount float64) {
	return rental.movie.determineAmount(rental.daysRented)
}

func (rental Rental) DeterminePoints() int {
	return rental.movie.determinePoints(rental.daysRented)
}

func (rental Rental) String() string {
	return fmt.Sprint(rental.movie)
}
