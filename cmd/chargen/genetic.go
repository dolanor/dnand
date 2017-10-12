package main

import (
	"math/rand"

	"github.com/MaxHalford/gago"
)

func (c Character) Evaluate() float64 {
	return float64(30 - (c.Dexterity.Bonus() + c.Strength.Bonus() + c.Constitution.Bonus() + c.HP))
}

func (c *Character) Mutate(rng *rand.Rand) {
	rate := 0.8
	if rng.Float64() > rate {
		return
	}
	ab := []*Ability{&c.Strength, &c.Dexterity, &c.Constitution}
	for i := range ab {
		mut := *ab[i] + Ability(*ab[i]*Ability(rng.NormFloat64()))
		if mut < 1 {
			mut = 1
		} else {
			mut = 18
		}
		*ab[i] = mut
	}
}

func (c *Character) Crossover(Y gago.Genome, rng *rand.Rand) (gago.Genome, gago.Genome) {
	return c, Y.(*Character)
}

func (c *Character) Clone() gago.Genome {
	y := Character{
		Strength:     c.Strength,
		Dexterity:    c.Dexterity,
		Constitution: c.Constitution,
		HP:           c.HP,
	}

	return &y
}

func CharacterFactory(rng *rand.Rand) gago.Genome {
	//TODO use class HP die
	dice := []int{4, 6, 8, 10, 12}
	d := rand.Intn(len(dice))
	cn := Character{
		Strength:     Ability(d6(3)),
		Dexterity:    Ability(d6(3)),
		Constitution: Ability(d6(3)),
		HP:           roll(dice[d], 1),
	}

	return &cn
}
