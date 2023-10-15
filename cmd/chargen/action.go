package main

import (
//	"fmt"
)

type Entity interface {
	Character() Character
}

type Target interface {
	Entity
}

type Effects []Effect
type Effect struct {
}

type Actions []Action

type Action interface {
	ID() string
	Act() Effects
}

type Attack struct {
	src, tgt Entity
}

/*
func (a *Attack) Act(src Entity, tgt Target) Effects {
	src.Character().Attack(tgt.Character())
	return Effects{}
}*/
func (a *Attack) Act() Effects {
	a.src.Character().Attack(a.tgt.Character())
	return Effects{}
}

func (c *Character) Act() Effects {
	//TODO choose an action from its available action
	action := Attack{}
	return action.Act()
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

type NewActor interface {
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
	return Effects{}
}
