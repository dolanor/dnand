package main

import (
	"math"
)

type Ability int8

// Index of abilities
const (
	Str int = iota
	Dex
	Con
	Int
	Wis
	Cha
)

func (c Ability) Bonus() int {
	return int(math.Floor(float64(c-10) / 2))

}
