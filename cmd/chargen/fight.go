package main

import (
//	"fmt"
)

func RoundOrder(c1, c2 Character) []Character {
	i1 := d20(1) + c1.Dexterity.Bonus()
	i2 := d20(1) + c2.Dexterity.Bonus()
	if i1 > i2 {
		return []Character{c1, c2}
	} else {
		return []Character{c2, c1}
	}
}

func Fight(c1, c2 Character) (Character, int) {
	ro := RoundOrder(c1, c2)

	i := 1
	for ; ; i++ {
		//		fmt.Printf("Round %d\n", i)
		for o, c := range ro {
			//			fmt.Printf("%s's turn… [%d HP]\n", c.Name, c.HP)
			var tgt *Character
			if o == 0 {
				tgt = &ro[1]
			} else {
				tgt = &ro[0]
			}
			*tgt = c.Attack(*tgt)
			if tgt.HP < 1 {
				tgt.Lost++
				c.Won++
				fmt.Printf("FIGHT: %s Won!\n", c)
				return c, i
			}
		}
	}
}
