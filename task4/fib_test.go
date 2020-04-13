package task4

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	testCases := []struct {
		name, expected string
		length int
	}{
		{
			name: "five nums",
			expected: "1 1 2 3 5 ",
			length: 5,
		},
		{
			name: "eleven nums",
			expected: "1 1 2 3 5 8 13 21 34 55 89 ",
			length: 11,
		},
	}

	for _, tc := range testCases {
		f := fibonacci()
		t.Run(tc.name, func(t *testing.T) {
			actual := ""
			for i := 0; i < tc.length; i++ {
				actual += strconv.Itoa(f()) + " "
			}
			assert.EqualValues(t, tc.expected, actual)
		})
	}
}
