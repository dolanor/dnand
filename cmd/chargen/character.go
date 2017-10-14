package main

import (
	"fmt"

	"github.com/MaxHalford/gago"
)

type Characteristics struct {
	FirstName string
	LastName  string

	Strength     Ability
	Dexterity    Ability
	Constitution Ability
	Intelligence Ability
	Wisdom       Ability
	Charisma     Ability

	HP int
}
type Char struct {
	OriginalCharacteristics Characteristics
	Characteristics
}

type Character struct {
	Char
	GA *gago.GA
}

func (c Character) AC() int {
	return 10 + c.Dexterity.Bonus()
}
func (c Character) FullName() string {
	return c.FirstName + " " + c.LastName
}

func (c Character) String() string {
	return fmt.Sprintf("<%s> 💗 %d/%d 💪 %d 🏃 %d 🏠 %d", c.FullName(), c.HP, c.OriginalCharacteristics.HP, c.Strength, c.Dexterity, c.Constitution)
}
