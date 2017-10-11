package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func d4(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += rand.Intn(4) + 1
	}
	return sum
}

func d10(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += rand.Intn(10) + 1
	}
	return sum
}

func d20(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += rand.Intn(20) + 1
	}
	return sum
}
