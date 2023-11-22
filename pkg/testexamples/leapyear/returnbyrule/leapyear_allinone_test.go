package returnbyrule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsLeapYearAllCases(t *testing.T) {
	testCases := map[string]struct {
		year        int
		expected    bool
		expectedErr error
	}{
		"Leap year if divisible by 4 but not by 100": {
			year:        1996,
			expected:    true,
			expectedErr: nil,
		},
		"Leap year if divisible by 400": {
			year:        2000,
			expected:    true,
			expectedErr: nil,
		},
		"Not leap year if divisible by 4 and 100, but not 400": {
			year:        2100,
			expected:    false,
			expectedErr: nil,
		},
		"Not divisible by 4": {
			year:        2022,
			expected:    false,
			expectedErr: nil,
		},
		"Negative year is invalid": {
			year:        -1,
			expected:    false,
			expectedErr: fmt.Errorf("invalid year %q because it is not greater than 0", -1),
		},
		"Zero year is invalid ": {
			year:        0,
			expected:    false,
			expectedErr: fmt.Errorf("invalid year %q because it is not greater than 0", 0),
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			result, err := IsLeapYear(tc.year)

			// A test case shouldn't have conditional flow.
			if err != nil {
				if tc.expectedErr == nil {
					require.NoError(t, err, "Unexpected error", err)
				} else {
					require.EqualError(t, err, tc.expectedErr.Error(), "Expected error")
				}
			} else {
				require.Equal(t, tc.expected, result, "Expected result")
			}
		})
	}
}
