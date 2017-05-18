package videostore

import "fmt"

type RentalStatement struct {
	name   string
	total  float64
	points int
	body   string
}

func NewStatement(name string) *RentalStatement {
	return &RentalStatement{name: name}
}

func (statement *RentalStatement) makeHeader() string {
	return "Rental Record for " + statement.name + "\n"
}

func (statement *RentalStatement) Include(rental *Rental) {
	amount := rental.DetermineAmount()
	statement.total += amount
	statement.points += rental.DeterminePoints()
	statement.body += formatRental(rental, amount)
}

func formatRental(rental *Rental, amount float64) string {
	return fmt.Sprintf("\t%s\t$%.2f\n", rental, amount)
}

func (statement *RentalStatement) FormatStatement() string {
	return statement.makeHeader() + statement.body + statement.makeFooter()
}

func (statement *RentalStatement) makeFooter() string {
	return fmt.Sprintf("You owed $%.2f\n"+
		"You earned %d frequent renter points\n",
		statement.total, statement.points)
}

func (statement *RentalStatement) AmountOwed() float64 {
	return statement.total
}

func (statement *RentalStatement) PointsEarned() int {
	return statement.points
}
