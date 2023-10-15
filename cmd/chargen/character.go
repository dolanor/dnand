package main

import (
	"fmt"

	"github.com/MaxHalford/gago"
)

type Characteristics struct {
	ID        uint64
	FirstName string
	LastName  string

	Strength     Ability
	Dexterity    Ability
	Constitution Ability
	Intelligence Ability
	Wisdom       Ability
	Charisma     Ability

	//Skills map[SkillID]Skill

	HP     int
	HPDice int
}

type SkillID uint8

const (
	Appraise SkillID = iota
	Balance
	Bluff
	Climb
	Concentration
	Craft
	DecipherScript
	Diplomacy
	DisableDevice
	Disguise
	EscapeArtist
	Forgery
	GatherInformation
	HandleAnimal
	Heal
	Hide
	Intimidate
	Jump
	Knowledge
	//TODO continue list
)

var SkillName = [...]string{
	"Appraise",
	"Balance",
	"Bluff",
	"Climb",
	"Concentration",
	"Craft",
	"DecipherScript",
	"Diplomacy",
	"DisableDevice",
	"Disguise",
	"EscapeArtist",
	"Forgery",
	"GatherInformation",
	"HandleAnimal",
	"Heal",
	"Hide",
	"Intimidate",
	"Jump",
	"Knowledge",
	//TODO continue list
}

type Skill struct {
	UserAbility *Ability
	Rank        int
	Desc        string
}

func (s Skill) Check(DC int) int {
	return d20(1) + s.Rank + s.UserAbility.Bonus() /* + Misc Modifiers*/ - DC
}

type Char struct {
	OriginalCharacteristics Characteristics
	Characteristics
}

type Character struct {
	Char
	GA               *gago.GA
	Won, Lost        int
	availableActions Actions
}

func (c Character) AvailableActions() Actions {
	return c.availableActions
}

func (c Character) AC() int {
	return 10 + c.Dexterity.Bonus()
}
func (c Character) FullName() string {
	return c.FirstName + " " + c.LastName
}

func (c Character) String() string {
	return fmt.Sprintf("<%s> 💗 %d/%d 💪 %d 🏃 %d 🏠 %d _-_ %d/%d", c.FullName(), c.HP, c.OriginalCharacteristics.HP, c.Strength, c.Dexterity, c.Constitution, c.Won, c.Lost)
}
