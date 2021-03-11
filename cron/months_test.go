package cron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseMonths(t *testing.T) {
	//c1 := "*"
	//c2 := "7"
	//c3 := "*/2"
	//c4 := "3,4,5"
	//c5 := "JAN,FEB"
	//c6 := "MAR,6,JUL"
	//c7 := "MAR, 6, JUM"
	//c8 := "JAM"
	//c9 := "13"
	
	cron := Cron{}

	testCases := []struct{
		name string
		input string
		expected []int
		expectedError error
	}{
		{
			name:          "asterisk only",
			input:         "*",
			expected:      monthsFullYear,
			expectedError: nil,
		},
		{
			name:          "asterisk slash 1",
			input:         "*/1",
			expected:      monthsFullYear,
			expectedError: nil,
		},
		{
			name: "number only",
			input: "7",
			expected: []int{7},
			expectedError: nil,
		},
		{
			name: "asterisk with division",
			input: "*/2",
			expected: []int{1,3,5,7,9,11},
			expectedError: nil,
		},
		{
			name: "asterisk with division by 3",
			input: "*/3",
			expected: []int{1,4,7,10},
			expectedError: nil,
		},
		{
			name: "asterisk with division by 4",
			input: "*/4",
			expected: []int{1,5,9},
			expectedError: nil,
		},
		{
			name: "comma separated",
			input: "3,4,5",
			expected: []int{3,4,5},
			expectedError: nil,
		},
		{
			name: "words check",
			input: "JAN,FEB",
			expected: []int{1,2},
			expectedError: nil,
		},
		{
			name: "words and numbers check",
			input: "JAN,6,JUL",
			expected: []int{1,6,7},
			expectedError: nil,
		},
		{
			name: "with space",
			input: "MAR, 6, JUL",
			expected: nil,
			expectedError: fmt.Errorf("cannot parse month: %s", `strconv.Atoi: parsing " 6": invalid syntax`),
		},
		{
			name: "wrong word",
			input: "JAM",
			expected: nil,
			expectedError: fmt.Errorf(`cannot parse month: strconv.Atoi: parsing "%s": invalid syntax`, "JAM"),
		},
	}


	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := cron.parseMonths(tc.input)
			require.Equal(t, tc.expected, cron.Months)
			assert.Equal(t, tc.expectedError, err)
		})

	}
}
