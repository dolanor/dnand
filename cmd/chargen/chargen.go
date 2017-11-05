package main

import (
	"fmt"

	"github.com/MaxHalford/gago"
)

func main() {
	var ga gago.GA
	ga = gago.Generational(BestCharacterFactory(&ga))
	ga.Model = gago.ModGenerational{Selector: SelFight{NContestants: 3}, MutRate: 0.5}
	ga.Initialize()

	fmt.Printf("Best fitness at generation 0: %f. %v\n", ga.Best.Fitness, ga.Best.Genome, len(ga.Populations[0].Individuals), ga.Populations[0].Individuals)
	for i := 1; i < 100; i++ {
		ga.Enhance()
		fmt.Printf("Best fitness at generation %d: %f. %v\n", i, ga.Best.Fitness, ga.Best.Genome)
	}
	fmt.Println("Populations:", len(ga.Populations[0].Individuals), ga.Populations[0].Individuals, ga)
}
