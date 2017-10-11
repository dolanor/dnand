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
	HP                                int
}

func (c Character) AC() int {
	return 10 + c.Dexterity.Bonus()
}

func (c Character) Attack(o Character) Character {
	roll := d20(1)
	bab := c.Strength.Bonus()
	if roll+bab < o.AC() {
		fmt.Printf("%s didn't get hit\n", o.Name)
		return o
	}

	dmg := d4(1)
	o.HP -= dmg
	fmt.Printf("%s lost %d HP\n", o.Name, dmg)
	return o
}

func RoundOrder(c1, c2 Character) []Character {
	i1 := d20(1) + c1.Dexterity.Bonus()
	i2 := d20(1) + c2.Dexterity.Bonus()
	if i1 > i2 {
		return []Character{c1, c2}
	} else {
		return []Character{c2, c1}
	}
}

func Fight(c1, c2 Character) Character {
	ro := RoundOrder(c1, c2)

	i := 1
	for ; ; i++ {
		fmt.Printf("Round %d\n", i)
		for o, c := range ro {
			fmt.Printf("%s's turn… [%d HP]\n", c.Name, c.HP)
			var tgt *Character
			if o == 0 {
				tgt = &ro[1]
			} else {
				tgt = &ro[0]
			}
			*tgt = c.Attack(*tgt)
			if tgt.HP < 1 {
				return c
			}
		}
	}
}

func main() {
	fmt.Println(d20(3))
	winner := Fight(Character{Name: "Corum", HP: d10(1), Strength: 10, Dexterity: 10, Constitution: 10},
		Character{Name: "Dwalin", HP: d10(1), Strength: 12, Dexterity: 8, Constitution: 12})
	fmt.Println("Winner is", winner.Name)
}
