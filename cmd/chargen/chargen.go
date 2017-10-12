package main

import (
	"fmt"

	"github.com/MaxHalford/gago"
)

func main() {
	winner := Fight(Character{Name: "Corum", HP: d10(1), Strength: 10, Dexterity: 10, Constitution: 10},
		Character{Name: "Dwalin", HP: d10(1), Strength: 12, Dexterity: 8, Constitution: 12})
	fmt.Println("Winner is", winner.Name)
	var ga = gago.Generational(CharacterFactory)
	ga.Initialize()

	fmt.Printf("Best fitness at generation 0: %f. %v\n", ga.CurrentBest.Fitness, ga.CurrentBest.Genome)
	for i := 1; i < 10; i++ {
		ga.Enhance()
		fmt.Printf("Best fitness at generation %d: %f. %v\n", i, ga.CurrentBest.Fitness, ga.CurrentBest.Genome)
	}
}
