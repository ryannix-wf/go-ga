package genetic

import (
	"math"
	"math/rand"
)

type Mapper interface {
	Distance(a int, b int) float64
	TotalDistance(path []int) int64
	Size() int
	Closest(int) int
}

type Location struct {
	x, y float64
}

type EuclideanMap struct {
	position    []Location
	distanceMap [][]float64
}

func (m *EuclideanMap) Distance(cityA int, cityB int) float64 {
	x := math.Abs(m.position[cityA].x - m.position[cityB].x)
	y := math.Abs(m.position[cityA].y - m.position[cityB].y)
	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	return dist
}

func (e *EuclideanMap) TotalDistance(path []int) int64 {
	var sum int64 = 0
	for index, element := range path {
		if index == (len(path) - 1) { //0 offset?
			sum = sum + int64(e.Distance(element, path[0]))
		} else {
			sum = sum + int64(e.Distance(element, path[index+1]))
		}
	}
	return sum
}

func (e *EuclideanMap) Size() int {
	return len(e.position)
}

func (e *EuclideanMap) Closest(pointA int) int {
	group := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		group[i] = rand.Intn(e.Size())
	}

	bestDist := 0.0
	bestLoc := 0
	for _, city := range group {
		dist := e.Distance(pointA, city)
		if bestDist == 0 || dist < bestDist {
			bestDist = dist
			bestLoc = city
		}
	}
	return bestLoc
}
