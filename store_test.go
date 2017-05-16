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

func (this *CustomerFixture) Setup() {
	this.customer = &Customer{"Fred", nil}
}

func (this *CustomerFixture) TestSingleNewReleaseStatement() {
	this.customer.Add(&Rental{&Movie{"The Cell", NewRelease}, 3})

	this.assertStatement("Rental Record for Fred\n" +
		"\tThe Cell\t9\n" +
		"You owed 9\n" +
		"You earned 2 frequent renter points\n")
}

func (this *CustomerFixture) TestDualReleaseStatement() {
	this.customer.Add(&Rental{&Movie{"The Cell", NewRelease}, 3})
	this.customer.Add(&Rental{&Movie{"The Tigger Movie", NewRelease}, 3})

	this.assertStatement("Rental Record for Fred\n" +
		"\tThe Cell\t9\n" +
		"\tThe Tigger Movie\t9\n" +
		"You owed 18\n" +
		"You earned 4 frequent renter points\n")
}

func (this *CustomerFixture) TestSingleChildrensStatement() {
	this.customer.Add(&Rental{&Movie{"The Tigger Movie", ChildrensMovie}, 3})

	this.assertStatement("Rental Record for Fred\n" +
		"\tThe Tigger Movie\t1.5\n" +
		"You owed 1.5\n" +
		"You earned 1 frequent renter points\n")
}

func (this *CustomerFixture) TestMultipleRegularStatement() {
	this.customer.Add(&Rental{&Movie{"Plan 9 from Outer Space", RegularMovie}, 1})
	this.customer.Add(&Rental{&Movie{"8 1/2", RegularMovie}, 2})
	this.customer.Add(&Rental{&Movie{"Eraserhead", RegularMovie}, 3})

	this.assertStatement("Rental Record for Fred\n" +
		"\tPlan 9 from Outer Space\t2\n" +
		"\t8 1/2\t2\n" +
		"\tEraserhead\t3.5\n" +
		"You owed 7.5\n" +
		"You earned 3 frequent renter points\n")
}

func (this *CustomerFixture) assertStatement(expected interface{}) {
	this.So(this.customer.Statement(), should.Equal, expected)
}
