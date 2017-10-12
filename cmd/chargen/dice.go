package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func roll(v, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += rand.Intn(v) + 1
	}
	return sum
}

func d4(n int) int {
	return roll(4, n)
}

func d6(n int) int {
	return roll(6, n)
}

func d8(n int) int {
	return roll(8, n)
}

func d10(n int) int {
	return roll(10, n)
}

func d12(n int) int {
	return roll(12, n)
}

func d20(n int) int {
	return roll(20, n)
}
