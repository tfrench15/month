package month

import "time"

// Month contains the current month.
type Month struct {
	current string
	next    string
}

// New returns a new instance of a Month.
func New() *Month {
	_, month, _ := time.Now().Date()

	cur := month.String()
	var nxt time.Month

	switch month {
	case time.January:
		nxt = time.February
	case time.February:
		nxt = time.March
	case time.March:
		nxt = time.April
	case time.April:
		nxt = time.May
	case time.May:
		nxt = time.June
	case time.June:
		nxt = time.July
	case time.July:
		nxt = time.August
	case time.August:
		nxt = time.September
	case time.September:
		nxt = time.October
	case time.October:
		nxt = time.November
	case time.November:
		nxt = time.December
	case time.December:
		nxt = time.January
	}

	return &Month{
		current: cur,
		next:    nxt.String(),
	}
}

// Show returns the current and month and is passed into the Lambda function.
func (m *Month) Show() (*Month, error) {
	return m, nil
}
