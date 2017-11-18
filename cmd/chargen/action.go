package main

import (
//	"fmt"
)

type Entity struct {
	Target
}

type Target interface {
	Character() Character
}

type Effects []Effect
type Effect struct {
}

type Actionable interface {
	Act(src Entity, tgt Target) Effects
}

type Attack struct {
}

func (a *Attack) Act(src Entity, tgt Target) Effects {
	src.Character().Attack(tgt.Character())
	return Effects{}
}

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

type NewActionable interface {
	Act() Effects
}

type MultiAttack struct {
	Src     Entity
	Targets []Target
}

// Usage:
// ma:= MultiAttack{Src: Lala, Targets: []Target{George, Pierre, Paul}}
// ma.Act()
func (ma *MultiAttack) Act() Effects {
	for _, t := range ma.Targets {
		ma.Src.Character().Attack(t.Character())
	}
}
