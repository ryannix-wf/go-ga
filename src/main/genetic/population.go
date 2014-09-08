package genetic

import (
	"math/rand"
)

type Population struct {
	individuals []Individual //THIS IS SORTED
	mapper      Mapper
}

func NewPopulation(size int, mapper Mapper) *Population {
	individuals := make([]Individual, 0, size) //maximum size of the population is size
	for i := 0; i < size; i++ {
		individuals = insertIndividual(individuals, GenerateIndividual(mapper))
	}
	pop := Population{individuals, mapper}

	return &pop
}

/*
Inserts the individual in an order from best fitness to worst.
*/
func insertIndividual(pop []Individual, ind *Individual) []Individual {
	//no longer needing to keep this ordered
	pop = append(pop, *ind)
	return pop

	/*
	   code to keep order in array on insert
	*/
	//lowerbound := 0
	//upperbound := len(pop)
	//count := 0
	//for lowerbound != upperbound {
	//	if count > 5 {
	//		break
	//	}
	//	count = count + 1
	//	halfway := int((len(pop[lowerbound:upperbound]) / 2)) + lowerbound
	//	if ind.distance < pop[halfway].distance {
	//		upperbound = halfway
	//	} else {
	//		lowerbound = halfway + 1 //offset to move into higher half
	//	}
	//}
	////we will not worry about capacity because the size is known and set.
	//pop = pop[:len(pop)+1]
	//copy(pop[lowerbound+1:], pop[lowerbound:])
	//pop[lowerbound] = *ind
	//return pop
}

func NextGeneration(oldPop *Population, mutationRate float64, mapper Mapper) *Population {
	newPop := make([]Individual, 0, len(oldPop.individuals))

	//keep the top 10% fit individuals
	//
	newPop = append(newPop, selectTopIndividuals(oldPop.individuals, 0.1)...)
	//topFit := int(len(oldPop.individuals) / 10.0)
	//copy(newPop, oldPop.individuals[:topFit]) //keep top 10
	//slice := oldPop.individuals[topFit:]

	created := make(chan *Individual, 2)
	signal := make(chan bool)
	go tournamentRunner(oldPop.individuals, created, signal)
	for len(newPop) < len(oldPop.individuals) {
		mom := <-created
		dad := <-created
		newPop = insertIndividual(newPop, SpawnChild(mom, dad, mutationRate, mapper))
	}
	signal <- true
	close(signal)
	close(created)
	return &Population{newPop, oldPop.mapper}
}

func selectTopIndividuals(pop []Individual, percentage float64) []Individual {
	top := make([]Individual, 0, int(percentage*float64(len(pop))))
	remaining := make([]Individual, len(pop), cap(pop))
	copy(remaining, pop)

	for i := 0; i < cap(top); i++ {
		best := 0
		bestfit := pop[0].distance
		for j := 1; j < len(pop); j++ {
			if pop[i].distance < bestfit {
				best = j
				bestfit = pop[i].distance]
			}
		}
		top = append(top, best)
	}
	return top
}

//runs tournament, kicks back winners
func tournamentRunner(pop []Individual, c chan *Individual, signal chan bool) {
	dead := false
	for dead != true {
		select {
		case dead = <-signal:
		default:
			best := pop[rand.Intn(len(pop))]
			for i := 0; i < Selection; i++ {
				competitor := pop[rand.Intn(len(pop))]
				if competitor.distance < best.distance {
					best = competitor
				}
			}
			c <- &best
		}
	}
}

func tournamentSelect(pop []Individual) *Individual {
	best := pop[rand.Intn(len(pop))]
	for i := 0; i < Selection; i++ {
		competator := pop[rand.Intn(len(pop))]
		if competator.distance < best.distance {
			best = competator
		}
	}
	return &best
}

func crossover(mother *Individual, father *Individual) []int {
	/*
		select a crossover point, correct dups
		mutate
		return new individual
	*/
	crossoverpoint := rand.Intn(len(mother.path))
	return append(mother.path[:crossoverpoint], father.path[crossoverpoint:]...)
}

func (pop *Population) BestDistance() int64 {
	return pop.individuals[0].distance
}
