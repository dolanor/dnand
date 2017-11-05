package main

import (
	"fmt"
	"math/rand"
)

// Utils functions mostly copypasted from gago
func newInts(n int) []int {
	var ints = make([]int, n)
	for i := range ints {
		ints[i] = i
	}
	return ints
}

// Sample k unique integers from a slice of n integers without replacement.
func sampleInts(ints []int, k int, rng *rand.Rand) ([]int, []int, error) {
	if k > len(ints) {
		return nil, nil, fmt.Errorf("Cannot sample %d elements from array of length %d", k, len(ints))
	}
	var (
		sample = make([]int, k)
		idxs   = make([]int, k)
	)
	for i, idx := range randomInts(k, 0, len(ints), rng) {
		sample[i] = ints[idx]
		idxs[i] = idx
	}
	return sample, idxs, nil
}

// Sample k unique integers in range [min, max) using reservoir sampling,
// specifically Vitter's Algorithm R.
func randomInts(k, min, max int, rng *rand.Rand) (ints []int) {
	ints = make([]int, k)
	for i := 0; i < k; i++ {
		ints[i] = i + min
	}
	for i := k; i < max-min; i++ {
		var j = rng.Intn(i + 1)
		if j < k {
			ints[j] = i + min
		}
	}
	return
}
