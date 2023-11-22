package singlereturn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLeapYearSingleReturn(t *testing.T) {
	require.True(t, IsLeapYear(2000))
}

func TestIsLeapYear(t *testing.T) {
	// Test cases for leap years
	leapYears := []int{2000, 2004, 2008, 2012, 2016, 2020}
	for _, year := range leapYears {
		require.True(t, IsLeapYear(year), "%d should be a leap year", year)

	}

	// Test cases for non-leap years
	nonLeapYears := []int{1900, 2001, 2002, 2003, 2005, 2010}
	for _, year := range nonLeapYears {
		require.False(t, IsLeapYear(year), "%d should not be a leap year", year)
	}
}
