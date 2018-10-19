// Package leap deals with logic surrounding leap years
package leap

// IsLeapYear takes a year and returns true or false depending on whether it is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%400 == 0 || year%100 != 0)
}
