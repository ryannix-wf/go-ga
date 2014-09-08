package genetic

import (
	"math/rand"
)

type Individual struct {
	path     []int
	distance int64
}

func GenerateIndividual(mapper Mapper) *Individual {
	size := mapper.Size()
	path := make([]int, size, size)
	cities := make([]int, size, size)
	//seems odd to be doing this
	for i := 0; i < size; i++ {
		cities[i] = i
	}

	for i := 0; i < size; i++ {
		x := rand.Intn(len(cities))
		path[i] = cities[x]
		if x == 0 {
			cities = cities[1:]
		} else if x == len(cities) {
			cities = cities[0 : x-1]
		} else {
			cities = append(cities[0:x], cities[x+1:]...)
		}
	}

	/*
	   Code to provide a better start location
	*/
	//start := rand.Intn(len(cities))
	//path[0] = start
	//end := 0
	//for i := 1; i < size; i++ {
	//	end = mapper.Closest(start)
	//	path[i] = end
	//	start = end
	//}

	return buildIndividual(path, 0, mapper) //path should be a slice right here
}

func randomShortestPath(mapper Mapper) []int {
	size := mapper.Size()
	startPoint := rand.Intn(size)
	path := make([]int, 0, size)
	path = append(path, startPoint)
	for i := 0; i < size-1; i++ {
		path = append(path, mapper.Closest(path[i]))
	}
	return path
}

func buildIndividual(path []int, mutationRate float64, mapper Mapper) *Individual {
	mutate(path, mutationRate)
	fixPath(path)
	dist := mapper.TotalDistance(path)
	return &Individual{path, dist}
}

func SpawnChild(mother *Individual, father *Individual, mutationRate float64, mapper Mapper) *Individual {
	/*
		select a crossover point, correct dups
		mutate
		return new individual
	*/
	crossoverpoint := rand.Intn(len(mother.path))
	i := buildIndividual(append(mother.path[:crossoverpoint], father.path[crossoverpoint:]...), MutationRate, mapper)
	return i
}

func mutate(path []int, mutationRate float64) {
	if mutationRate > 0 {
		for index, elem := range path {
			r := rand.Float64()
			if r <= mutationRate {
				//flip this location with another random location
				loc := rand.Intn(len(path))
				path[index] = path[loc]
				path[loc] = elem
			}
		}
	}
}

func fixPath(path []int) {
	/*
		traverse the path, record the city indices. If a duplicate is found:
			record location
		For every location of dups, reassign to the known unvisited cities
	*/
	unvisited := make(map[int]bool) //key is city
	dups := make(map[int]bool)      //key is location
	//initialize set to false on everything?
	for city, _ := range path {
		unvisited[city] = true
	}

	for loc, city := range path {
		if !unvisited[city] {
			dups[loc] = true
		}
		delete(unvisited, city)
	}

	for loc, _ := range dups {
	Reassignment:
		for city, _ := range unvisited {
			path[loc] = city
			delete(unvisited, city)
			break Reassignment //remember to break out!
		}
	}
}
