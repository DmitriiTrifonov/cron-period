package cron

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHoursParse(t *testing.T) {
	cron := Cron{}

	everyHour := make([]int, 24)

	for i := range everyHour {
		everyHour[i] = i
	}

	testCases := []struct {
		name          string
		input         string
		expected      []int
		expectedError error
	}{
		{
			name:          "asterisk only",
			input:         "*",
			expected:      everyHour,
			expectedError: nil,
		},
		{
			name:          "asterisk slash 1",
			input:         "*/1",
			expected:      everyHour,
			expectedError: nil,
		},
		{
			name:          "number only",
			input:         "7",
			expected:      []int{7},
			expectedError: nil,
		},
		{
			name:          "asterisk with division",
			input:         "*/2",
			expected:      []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22},
			expectedError: nil,
		},
		{
			name:          "asterisk with division by 3",
			input:         "*/3",
			expected:      []int{0, 3, 6, 9, 12, 15, 18, 21},
			expectedError: nil,
		},
		{
			name:          "asterisk with division by 4",
			input:         "*/4",
			expected:      []int{0, 4, 8, 12, 16, 20},
			expectedError: nil,
		},
		{
			name:          "comma separated",
			input:         "3,4,5",
			expected:      []int{3, 4, 5},
			expectedError: nil,
		},
		{
			name:          "word",
			input:         "HOUR",
			expected:      nil,
			expectedError: fmt.Errorf("cannot parse hour: %s", `strconv.Atoi: parsing "HOUR": invalid syntax`),
		},
		{
			name:          "words with comma",
			input:         "ANOTHER,WORD",
			expected:      nil,
			expectedError: fmt.Errorf(`cannot parse hour: strconv.Atoi: parsing "%s": invalid syntax`, "ANOTHER"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cron.parseHours(tc.input)
			require.Equal(t, tc.expected, cron.Hours)
			assert.Equal(t, tc.expectedError, err)
		})

	}
}
