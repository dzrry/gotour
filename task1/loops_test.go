package task1

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	const delta = 0.5

	testCases := []struct {
		name          string
		inputValue    float64
		expectedValue float64
	}{
		{
			name:          "two",
			inputValue:    2,
			expectedValue: 1.4,
		},
		{
			name:          "four",
			inputValue:    4,
			expectedValue: 2,
		},
		{
			name:          "twentyfive",
			inputValue:    25,
			expectedValue: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualValue := Sqrt(tc.inputValue)
			diff := math.Abs(tc.expectedValue - actualValue)
			if diff > delta {
				t.Errorf("Actual is not in delta range: \n"+
					"expected: %v\n"+
					"actual: %v\n"+
					"diff: %v", tc.expectedValue, actualValue, diff)
			}
		})
	}
}
