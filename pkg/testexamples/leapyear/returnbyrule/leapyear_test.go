package returnbyrule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLeapYear(t *testing.T) {
	testCases := map[string]struct {
		year     int
		expected bool
	}{
		"Leap year if divisible by 4 but not by 100": {
			year:     1996,
			expected: true,
		},
		"Leap year if divisible by 400": {
			year:     2000,
			expected: true,
		},
		"Not leap year if divisible by 4 and 100, but not 400": {
			year:     1900,
			expected: false,
		},
		"Not divisible by 4": {
			year:     2022,
			expected: false,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			result, err := IsLeapYear(tc.year)
			require.NoErrorf(t, err, "year %d is a valid year. unexpected error: %s", tc.year, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestIsLeapYearError(t *testing.T) {
	testCases := map[string]struct {
		year     int
		expected error
	}{

		"Negative year is invalid": {
			year:     -1,
			expected: fmt.Errorf("invalid year %q because it is not greater than 0", -1),
		},
		"Zero year is invalid ": {
			year:     0,
			expected: fmt.Errorf("invalid year %q because it is not greater than 0", 0),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := IsLeapYear(tc.year)
			require.Error(t, err, "year %d is an invalid year. validation must fail")
		})
	}
}
