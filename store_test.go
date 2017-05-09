package videostore

import "testing"

func TestSingleNewReleaseStatement(t *testing.T) {
	customer := &Customer{"Fred", nil}
	customer.Add(&Rental{&Movie{"The Cell", NewRelease}, 3})

	want := "Rental Record for Fred\n\tThe Cell\t9\nYou owed 9\nYou earned 2 frequent renter points\n"
	if got := customer.Statement(); got != want {
		t.Errorf("Got: %s\nWant: %s", got, want)
	}
}

func TestDualReleaseStatement(t *testing.T) {
	customer := &Customer{"Fred", nil}
	customer.Add(&Rental{&Movie{"The Cell", NewRelease}, 3})
	customer.Add(&Rental{&Movie{"The Tigger Movie", NewRelease}, 3})

	want := "Rental Record for Fred\n\tThe Cell\t9\n\tThe Tigger Movie\t9\nYou owed 18\nYou earned 4 frequent renter points\n"
	if got := customer.Statement(); got != want {
		t.Errorf("Got: %s\nWant: %s", got, want)
	}
}

func TestSingleChildrensStatement(t *testing.T) {
	customer := &Customer{"Fred", nil}
	customer.Add(&Rental{&Movie{"The Tigger Movie", ChildrensMovie}, 3})

	want := "Rental Record for Fred\n\tThe Tigger Movie\t1.5\nYou owed 1.5\nYou earned 1 frequent renter points\n"
	if got := customer.Statement(); got != want {
		t.Errorf("Got: %s\nWant: %s", got, want)
	}
}

func TestMultipleRegularStatement(t *testing.T) {
	customer := &Customer{"Fred", nil}
	customer.Add(&Rental{&Movie{"Plan 9 from Outer Space", RegularMovie}, 1})
	customer.Add(&Rental{&Movie{"8 1/2", RegularMovie}, 2})
	customer.Add(&Rental{&Movie{"Eraserhead", RegularMovie}, 3})

	want := "Rental Record for Fred\n\tPlan 9 from Outer Space\t2\n\t8 1/2\t2\n\tEraserhead\t3.5\nYou owed 7.5\nYou earned 3 frequent renter points\n"
	if got := customer.Statement(); got != want {
		t.Errorf("Got: %s\nWant: %s", got, want)
	}
}
