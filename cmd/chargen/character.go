package main

import (
	"fmt"

	"github.com/MaxHalford/gago"
)

type Characteristics struct {
	Name                              string
	Strength, Dexterity, Constitution Ability
	HP                                int
}
type Character struct {
	OriginalCharacteristics Characteristics
	Characteristics
	GA *gago.GA
}

func (c Character) AC() int {
	return 10 + c.Dexterity.Bonus()
}

func (c Character) String() string {
	return fmt.Sprintf("<%s> 💗 %d/%d 💪 %d 🏃 %d 🏠 %d", c.Name, c.HP, c.OriginalCharacteristics.HP, c.Strength, c.Dexterity, c.Constitution)
}
