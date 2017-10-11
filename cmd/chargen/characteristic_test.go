package main

import (
	"testing"
)

func TestBonus(t *testing.T) {
	tab := []struct {
		val      int
		expected int
	}{
		{0, -5}, {1, -5}, {2, -4}, {3, -4}, {4, -3}, {5, -3},
		{6, -2}, {7, -2}, {8, -1}, {9, -1}, {10, 0}, {11, 0},
		{12, 1}, {13, 1}, {14, 2}, {15, 2}, {16, 3}, {17, 3},
		{18, 4}, {19, 4}, {20, 5}, {21, 5}, {22, 6},
	}

	for _, v := range tab {
		got := Ability(v.val).Bonus()
		if got != v.expected {
			t.Fatalf("Wrong bonus computation.\n\tgot = %v\n\texpected = %v", got, v.expected)
		}
	}

}
