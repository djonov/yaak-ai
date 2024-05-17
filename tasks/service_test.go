package tasks

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		a        int
		b        int
		expected int
	}{
		"2+5": {
			a:        2,
			b:        5,
			expected: 7,
		},
		"5+-2": {
			a:        5,
			b:        -2,
			expected: 3,
		},
		"-1+6": {
			a:        -1,
			b:        6,
			expected: 5,
		},
		"0+6": {
			a:        0,
			b:        6,
			expected: 6,
		},
		"1+0": {
			a:        1,
			b:        0,
			expected: 1,
		},
		"0+0": {
			a:        0,
			b:        0,
			expected: 0,
		},
		"-3+-4": {
			a:        -3,
			b:        -4,
			expected: -7,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Add(test.a, test.b)

			if actual != test.expected {
				t.Errorf("Sums do not match: expected %v, got %v", test.expected, actual)
			}
		})
	}
}
