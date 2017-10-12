package main

import (
	"fmt"

	"github.com/MaxHalford/gago"
)

func main() {
	winner := Fight(Character{Characteristics: Characteristics{Name: "Corum", HP: d10(1), Strength: 10, Dexterity: 10, Constitution: 10}},
		Character{Characteristics: Characteristics{Name: "Dwalin", HP: d10(1), Strength: 12, Dexterity: 8, Constitution: 12}})
	fmt.Println("Winner is", winner.Name)
	var ga gago.GA
	ga = gago.Generational(BestCharacterFactory(&ga))
	ga.Initialize()

	fmt.Printf("Best fitness at generation 0: %f. %v\n", ga.Best.Fitness, ga.Best.Genome, len(ga.Populations[0].Individuals), ga.Populations[0].Individuals)
	for i := 1; i < 100; i++ {
		ga.Enhance()
		fmt.Printf("Best fitness at generation %d: %f. %v\n", i, ga.Best.Fitness, ga.Best.Genome)
	}
	fmt.Println("Populations:", len(ga.Populations[0].Individuals), ga.Populations[0].Individuals, ga)
}
