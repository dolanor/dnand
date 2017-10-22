package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/Pallinder/sillyname-go"
	"github.com/pkg/errors"

	"github.com/MaxHalford/gago"
)

func (c Character) Evaluate() float64 {
	tgt, ok := c.GA.CurrentBest.Genome.(*Character)
	if !ok {
		fmt.Println("Couldn't get the current best!")
		return 1000.0
	}
	start := time.Now()
	winner := Fight(c, *tgt)
	elapsed := time.Since(start)
	if c != winner {
		return 1000.0
	}
	return float64(winner.HP) / float64(winner.OriginalCharacteristics.HP) * float64(elapsed.Nanoseconds())

	//return float64(30 - (c.Dexterity.Bonus() + c.Strength.Bonus() + c.Constitution.Bonus() + c.HP))
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
	// We mostly crossover lastname and regenerate a new first name to track heritage
	n := c.Clone().(*Character)
	n.Char.FirstName = randomdata.FirstName(randomdata.RandomGender)

	return n, Y.(*Character)
}

func (c *Character) Clone() gago.Genome {
	y := Character{
		Char: Char{
			Characteristics:         c.Characteristics,
			OriginalCharacteristics: c.OriginalCharacteristics,
		},
		GA: c.GA,
	}

	return &y
}
func BestCharacterFactory(ga *gago.GA) func(*rand.Rand) gago.Genome {
	return func(rng *rand.Rand) gago.Genome {
		//TODO use class HP die
		dice := []int{4, 6, 8, 10, 12}
		d := rand.Intn(len(dice))
		fname := strings.Split(sillyname.GenerateStupidName(), " ")
		ch := Characteristics{
			FirstName:    fname[0],
			LastName:     fname[1],
			Strength:     Ability(d6(3)),
			Dexterity:    Ability(d6(3)),
			Constitution: Ability(d6(3)),
			HP:           roll(dice[d], 1),
		}
		cn := Character{
			Char: Char{
				OriginalCharacteristics: ch,
				Characteristics:         ch,
			},
			GA: ga,
		}

		return &cn
	}
}

type SelFight struct {
	NContestants int
}

// Apply SelTournament.
func (sel SelFight) Apply(n int, indis gago.Individuals, rng *rand.Rand) (gago.Individuals, []int, error) {
	// Check that the number of individuals is large enough
	if len(indis)-n < sel.NContestants-1 {
		return nil, nil, fmt.Errorf("Not enough individuals to select %d "+
			"with NContestants = %d, have %d individuals and need at least %d",
			n, sel.NContestants, len(indis), sel.NContestants+n-1)
	}
	var (
		winners         = make(gago.Individuals, n)
		indexes         = make([]int, n)
		notSelectedIdxs = newInts(len(indis))
	)
	for i := range winners {
		// Sample contestants
		var (
			contestants, idxs, _ = sampleInts(notSelectedIdxs, sel.NContestants, rng)
			winnerIdx            int
		)
		// Find the best contestant
		winners[i] = indis[contestants[0]]
		//winners[i].Evaluate()
		for j, idx := range contestants[1:] {
			indis[idx]
			if indis[idx].GetFitness() < winners[i].Fitness {
				winners[i] = indis[idx]
				indexes[i] = idx
				winnerIdx = idxs[j]
			}
		}
		// Ban the winner from re-participating
		notSelectedIdxs = append(notSelectedIdxs[:winnerIdx], notSelectedIdxs[winnerIdx+1:]...)
	}
	return winners.Clone(rng), indexes, nil
}

// Validate SelTournament fields.
func (sel SelFight) Validate() error {
	if sel.NContestants < 1 {
		return errors.New("NContestants should be higher than 0")
	}
	return nil
}
