package returnbyrule

import "fmt"

func IsLeapYear(year int) (bool, error) {
	// A year is valid if it is greater than 0
	if year <= 0 {
		return false, fmt.Errorf("invalid year %q because it is not greater than 0", year)
	}
	// A year is a leap year if it is divisible by 4 but not by 100
	if year%4 == 0 && year%100 != 0 {
		return true, nil
	}
	// A year is a leap year if it is divisible by 400
	if year%400 == 0 {
		return true, nil
	}
	// Otherwise it is not a leap year:
	// if it is not divisible by 4
	// if it is not a leap year if it is divisible by 4 and 100, but not 400
	return false, nil
}
