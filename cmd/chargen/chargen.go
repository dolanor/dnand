package main

import (
	"fmt"
)

func main() {
	winner := Fight(Character{Name: "Corum", HP: d10(1), Strength: 10, Dexterity: 10, Constitution: 10},
		Character{Name: "Dwalin", HP: d10(1), Strength: 12, Dexterity: 8, Constitution: 12})
	fmt.Println("Winner is", winner.Name)
}
