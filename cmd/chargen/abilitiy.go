package main

import (
	"math"
)

type Ability int8

func (c Ability) Bonus() int {
	return int(math.Floor(float64(c-10) / 2))

}
