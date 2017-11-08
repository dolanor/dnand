package main

import (
//	"fmt"
)

func (c Character) Attack(o Character) Character {
	hitRoll := d20(1)
	bab := c.Strength.Bonus()
	if hitRoll+bab < o.AC() {
		//		fmt.Printf("%s didn't get hit\n", o.Name)
		return o
	}

	nbDice := 1
	if hitRoll == 20 {
		nbDice = 2
	}
	dmg := roll(4, nbDice)
	o.HP -= dmg
	//	fmt.Printf("%s lost %d HP\n", o.Name, dmg)
	return o
}
