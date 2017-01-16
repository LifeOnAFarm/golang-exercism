/*
Name:		LifeOnAFarn
Date:		12/01/17
Program:	Program that accepts a year and tells you if it is a leap year
*/
package leap

const testVersion = 2

// Accepts a year and returns true or false
func IsLeapYear(year int) bool {

	// If the year can be divided by 4 or 400 evenly then it is a leap year,
	// unless it can also be divided by 100 evenly
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	} else {
		return false
	}
}
