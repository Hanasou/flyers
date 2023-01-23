package memdb

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinTableDriven(t *testing.T) {
	// tests is a list of structs containing a, b, and want
	// We load our test cases into here
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// Range over all our test cases
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
