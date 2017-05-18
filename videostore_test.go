package videostore

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestCustomerFixture(t *testing.T) {
	gunit.Run(new(CustomerFixture), t)
}

type CustomerFixture struct {
	*gunit.Fixture
	statement *RentalStatement
}

func (this *CustomerFixture) Setup() {
	this.statement = NewStatement("Customer Name")
}

func (this *CustomerFixture) assertOwedAndPoints(owed float64, points int) {
	this.So(this.statement.AmountOwed(), should.Equal, owed)
	this.So(this.statement.PointsEarned(), should.Equal, points)
}

func (this *CustomerFixture) assertStatement(expected string) {
	this.So(this.statement.FormatStatement(), should.Equal, expected)
}

func (this *CustomerFixture) TestSingleNewReleaseStatement() {
	this.statement.Include(NewRental(newRelease1, 3))
	this.assertOwedAndPoints(9.0, 2)
}

func (this *CustomerFixture) TestDualReleaseStatement() {
	this.statement.Include(NewRental(newRelease1, 3))
	this.statement.Include(NewRental(newRelease2, 3))
	this.assertOwedAndPoints(18, 4)
}

func (this *CustomerFixture) TestSingleChildrensStatement() {
	this.statement.Include(NewRental(childrens, 3))
	this.assertOwedAndPoints(1.5, 1)
}

func (this *CustomerFixture) TestMultipleRegularStatement() {
	this.statement.Include(NewRental(regular1, 1))
	this.statement.Include(NewRental(regular2, 2))
	this.statement.Include(NewRental(regular3, 3))
	this.assertOwedAndPoints(7.5, 3)
}

func (this *CustomerFixture) TestStatementFormatting() {
	this.statement.Include(NewRental(regular1, 1))
	this.statement.Include(NewRental(regular2, 2))
	this.statement.Include(NewRental(regular3, 3))

	this.assertStatement(
		"Rental Record for Customer Name\n" +
			"\tRegular 1\t$2.00\n" +
			"\tRegular 2\t$2.00\n" +
			"\tRegular 3\t$3.50\n" +
			"You owed $7.50\n" +
			"You earned 3 frequent renter points\n")
}

var (
	newRelease1 = &NewReleaseMovie{Movie{"New Release 1"}}
	newRelease2 = &NewReleaseMovie{Movie{"New Release 2"}}
	childrens   = &ChildrensMovie{Movie{"Childrens"}}
	regular1    = &Movie{"Regular 1"}
	regular2    = &Movie{"Regular 2"}
	regular3    = &Movie{"Regular 3"}
)
