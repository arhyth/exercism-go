// Leap exposes IsLeapYear function.
package leap

// IsLeapYear evaluates whether given int value is a leap year.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
