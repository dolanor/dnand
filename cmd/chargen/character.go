package main

type Character struct {
	Name                              string
	Strength, Dexterity, Constitution Ability
	HP                                int
}

func (c Character) AC() int {
	return 10 + c.Dexterity.Bonus()
}
