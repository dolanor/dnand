package main

import (
//	"fmt"
)

func (c Character) Attack(o Character) Character {
	roll := d20(1)
	bab := c.Strength.Bonus()
	if roll+bab < o.AC() {
		//		fmt.Printf("%s didn't get hit\n", o.Name)
		return o
	}

	dmg := d4(1)
	o.HP -= dmg
	//	fmt.Printf("%s lost %d HP\n", o.Name, dmg)
	return o
}
