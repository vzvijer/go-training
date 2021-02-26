// package leap implements functions related to determing of leap or not leap years
package leap

// IsLeapYear returns true if given year is a leap year
func IsLeapYear(y int) bool {
	return y%400 == 0 || (y%4 == 0 && y%100 != 0)
}
