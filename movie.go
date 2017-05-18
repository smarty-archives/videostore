package videostore

type (
	Movie           struct{ title string }
	ChildrensMovie  struct{ Movie }
	NewReleaseMovie struct{ Movie }
)

func (this Movie) String() string { return this.title }

/**************************************************************************/

func (Movie) determineAmount(daysRented int) (amount float64) {
	amount = 2
	if daysRented > 2 {
		amount += (float64)(daysRented-2) * 1.5
	}
	return amount
}
func (ChildrensMovie) determineAmount(daysRented int) (amount float64) {
	amount += 1.5
	if daysRented > 3 {
		amount += (float64)(daysRented-3) * 1.5
	}
	return amount
}
func (NewReleaseMovie) determineAmount(daysRented int) (amount float64) {
	return (float64)(daysRented * 3)
}

/**************************************************************************/

func (Movie) determinePoints(daysRented int) int          { return 1 }
func (ChildrensMovie) determinePoints(daysRented int) int { return 1 }
func (NewReleaseMovie) determinePoints(daysRented int) int {
	if daysRented > 1 {
		return 2
	} else {
		return 1
	}
}
