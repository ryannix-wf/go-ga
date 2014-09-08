package main

import (
	"log"
	"main/genetic"
)

func main() {

	mapper := genetic.LoadFromFile("uy734.tsp.txt")
	log.Println(mapper.Size())
	count := 0
	var f int64 = 0
	pop := genetic.NewPopulation(genetic.PopulationSize, mapper)
	bestFit := pop.BestDistance()
	log.Println(pop.BestDistance())
Rounds:
	for i := 0; i < genetic.Rounds; i++ {
		log.Println("Round ", i, " Stagnation: ", count)
		pop := genetic.NextGeneration(pop, genetic.MutationRate, mapper)
		f = pop.BestDistance()
		if f < bestFit || bestFit == 0.0 {
			bestFit = f
			count = 0
		} else if f >= bestFit {
			count = count + 1
		}
		if count == genetic.Stagnation {
			break Rounds
		}
		log.Println(bestFit)
	}
	log.Println("Best Route Distance is: ", bestFit)

}
