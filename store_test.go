package videostore

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestCustomerFixture(t *testing.T) { gunit.Run(new(CustomerFixture), t) }

type CustomerFixture struct {
	*gunit.Fixture

	customer *Customer
}

func (this *CustomerFixture) assertOwedAndPoints(owed float64, points int) {
	this.customer.Statement()
	this.So(this.customer.AmountOwed(), should.Equal, owed)
	this.So(this.customer.PointsEarned(), should.Equal, points)
}

func (this *CustomerFixture) assertStatement(expected interface{}) {
	this.So(this.customer.Statement(), should.Equal, expected)
}

func (this *CustomerFixture) Setup() {
	this.customer = &Customer{name: "Customer Name"}
}

func (this *CustomerFixture) TestSingleNewReleaseStatement() {
	this.customer.Add(&Rental{newRelease1, 3})
	this.assertOwedAndPoints(9.0, 2)
}

func (this *CustomerFixture) TestDualReleaseStatement() {
	this.customer.Add(&Rental{newRelease1, 3})
	this.customer.Add(&Rental{newRelease2, 3})
	this.assertOwedAndPoints(18, 4)
}

func (this *CustomerFixture) TestSingleChildrensStatement() {
	this.customer.Add(&Rental{childrens, 3})
	this.assertOwedAndPoints(1.5, 1)
}

func (this *CustomerFixture) TestMultipleRegularStatement() {
	this.customer.Add(&Rental{regular1, 1})
	this.customer.Add(&Rental{regular2, 2})
	this.customer.Add(&Rental{regular3, 3})
	this.assertOwedAndPoints(7.5, 3)
}

func (this *CustomerFixture) TestStatementFormatting() {
	this.customer.Add(&Rental{regular1, 1})
	this.customer.Add(&Rental{regular2, 2})
	this.customer.Add(&Rental{regular3, 3})

	this.assertStatement("Rental Record for Customer Name\n" +
		"\tRegular 1\t2\n" +
		"\tRegular 2\t2\n" +
		"\tRegular 3\t3.5\n" +
		"You owed 7.5\n" +
		"You earned 3 frequent renter points\n")
}
var (
	newRelease1 = &Movie{"New Release 1", NewRelease}
	newRelease2 = &Movie{"New Release 2", NewRelease}
	childrens   = &Movie{"Childrens", ChildrensMovie}
	regular1    = &Movie{"Regular 1", RegularMovie}
	regular2    = &Movie{"Regular 2", RegularMovie}
	regular3    = &Movie{"Regular 3", RegularMovie}
)
