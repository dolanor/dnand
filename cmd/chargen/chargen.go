package main

import (
	"fmt"
	"math"
)

type Ability int8

func (c Ability) Bonus() int {
	return int(math.Floor(float64(c-10) / 2))

}

type Character struct {
	Name                              string
	Strength, Dexterity, Constitution Ability
}

	}
}

func RoundInit(c1, c2 Character) []Character {
	i1 := d20(1) + c1.Dexterity.Bonus()
	i2 := d20(1) + c2.Dexterity.Bonus()
	if i1 > i2 {
		return []Character{c1, c2}
	} else {
		return []Character{c2, c1}
	}
}

func Fight(c1, c2 Character) Character {

	for {

		return c1
	}
}

func main() {
	fmt.Println(d20(3))
}
