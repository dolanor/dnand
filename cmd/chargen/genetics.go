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
	//tgt, ok := c.GA.CurrentBest.Genome.(*Character)
	//if !ok {
	//		fmt.Println("Couldn't get the current best!")
	//		return 1000.0
	//	}
	start := time.Now()
	//winner, _ := Fight(c, *tgt)
	elapsed := time.Since(start)
	//if c != winner {
	//	return 1000.0
	//}
	_ = elapsed
	//	return ((1.01 - (float64(winner.HP) / float64(winner.OriginalCharacteristics.HP))) * (0.01 * float64(elapsed.Nanoseconds())) * float64(rounds))
	tot := c.Won + c.Lost
	fmt.Printf("Evaluating %v\n", c)
	return float64(tot) - float64(c.Won)/float64(tot)

	// 10/10 * 1sec = 1 * 1
	// 9/10 * 1sec = 0.9
	// 1/10 * 3s = 0.3

	// 1 - 10/10 * 1s = 0 * 1s
	// 1 - 9/10 * 1s = 0.1 * 1 = .1
	// 1 - 1/10 * 3s = 0.9 * 3 = 2.7

	//return float64(30 - (c.Dexterity.Bonus() + c.Strength.Bonus() + c.Constitution.Bonus() + c.HP))
}

func (c *Character) Mutate(rng *rand.Rand) {
	rate := 0.8
	if rng.Float64() > rate {
		return
	}
	//HP
	dice := [...]int{4, 6, 8, 10, 12}
	d := rand.Intn(len(dice))
	mut := roll(dice[d], 1)
	c.Char.OriginalCharacteristics.HP = mut
	c.Char.HP = mut

	ab := [...]*Ability{&c.Strength, &c.Dexterity, &c.Constitution}
	for i := range ab {
		mut := *ab[i] + Ability(*ab[i]*Ability(rng.NormFloat64()))
		if mut == *ab[i] {
			continue
		}

		if mut < 1 {
			mut = 1
		} else if mut > 18 {
			mut = 18
		}

		// It is the constitution, we need to update HPs
		if i == Con {
			c.Char.OriginalCharacteristics.HP += -ab[Con].Bonus() + mut.Bonus()
			c.Char.HP += -ab[Con].Bonus() + mut.Bonus()
		}

		*ab[i] = mut
	}
}

func (c *Character) Crossover(Y gago.Genome, rng *rand.Rand) (gago.Genome, gago.Genome) {
	// We mostly crossover lastname and regenerate a new first name to track heritage
	n := c.Clone().(*Character)
	n.Char.FirstName = randomdata.FirstName(randomdata.RandomGender)
	n.ID = rand.Uint64()

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
		dice := [...]int{4, 6, 8, 10, 12}
		d := rand.Intn(len(dice))
		fname := strings.Split(sillyname.GenerateStupidName(), " ")
		ch := Characteristics{
			FirstName:    fname[0],
			LastName:     fname[1],
			Strength:     Ability(d6(3)),
			Dexterity:    Ability(d6(3)),
			Constitution: Ability(d6(3)),
			HPDice:       d,
		}
		for ; ch.HP <= 0; ch.HP = roll(dice[d], 1) + ch.Constitution.Bonus() {
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
			w, _ := Fight(*(winners[i].Genome.(*Character)), *(indis[idx].Genome.(*Character)))
			fmt.Printf("COMPARE %v VS %v\n", winners[i].Genome.(*Character), w)
			if w.ID != winners[i].Genome.(*Character).ID {
				winners[i].Genome.(*Character).Lost++
				indis[idx].Genome.(*Character).Won++
				winners[i] = indis[idx]
				indexes[i] = idx
				winnerIdx = idxs[j]
			} else {
				winners[i].Genome.(*Character).Won++
				indis[idx].Genome.(*Character).Lost++
			}
			curr := winners[i].Genome.(*Character)
			other := indis[idx].Genome.(*Character)
			fmt.Printf("Tournament: <%s> %d/%d VS <%s> %d/%d\n", curr.FullName(), curr.Won, curr.Lost, other.FullName(), other.Won, other.Lost)
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
