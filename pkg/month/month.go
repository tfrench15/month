package month

import "time"

// Get returns the current month.
func Get() string {
	_, month, _ := time.Now().Date()
	return month.String()
}
