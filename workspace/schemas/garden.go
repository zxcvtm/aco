package schemas

import (
	"math/rand"
	"time"
)

type (
	Garden struct {
 		graph     [][]Edge
 		antColony []Ant
 		alpha     float64
 		beta      float64
 		p         float64
	}
	Edge struct {
		Distance float64
		Pheromone float64
	}
)
func (garden Garden) InitGarden(graph [][]Edge, alpha, beta, pheromone, p float64) Garden {
	garden.alpha = alpha
	garden.beta = beta
	garden.p = p
	garden.graph = graph

	for x, edges := range garden.graph {
		for y, edge := range edges {
			edge.Pheromone = pheromone
			garden.graph[x][y] = edge
		}
	}
	return garden
}

func (garden Garden) Visibility(from, to int) float64 {
	return 1/garden.graph[from][to].Distance
}
func (garden Garden) GetPheromone(from, to int) float64 {
	return garden.graph[from][to].Pheromone
}
func (garden Garden) GetDistance(from, to int) float64 {
	return garden.graph[from][to].Distance
}

func (garden Garden) InitRandomAntColony () {
	for i :=0; i < len(garden.graph); i++ {
		ant := Ant{}
		rand.Seed(time.Now().Unix())
		randomStartNode := rand.Intn(len(garden.graph))
		ant = ant.InitAnt(randomStartNode, 1.0)
		garden.antColony = append(garden.antColony, )
	}
}
func (garden Garden) InitAntColony () Garden{
	for i :=0; i < len(garden.graph); i++ {
		ant := Ant{}
		ant = ant.InitAnt(0, 1.0)
		garden.antColony = append(garden.antColony, ant)
	}
	return garden
}
func (garden Garden) Iterate() Garden{
	garden = garden.InitAntColony()
	for i := 1; i < len(garden.graph); i++ {
		garden = garden.moveAnts()
	}
	garden = garden.antsAddPheromone()
	return garden
}
func (garden Garden) IterateAsync() Garden{
	garden = garden.InitAntColony()
	for i := 1; i < len(garden.graph); i++ {
		garden = garden.moveAntsAsync()
	}
	garden = garden.antsAddPheromone()
	return garden
}
func (garden Garden) moveAnts() (Garden){
	antColony := []Ant{}
	for _,ant := range garden.antColony {
		antColony = append(antColony, ant.Move(garden))
	}
	garden.antColony = antColony
	return  garden
}
func (garden Garden) moveAntsAsync() (Garden){
	var channelArray []<-chan Ant
	antColony := []Ant{}
	for _,ant := range garden.antColony {
		channelArray = append(channelArray,ant.MoveAsync(garden))
	}
	for _, channelResult := range channelArray {
		response := <-channelResult
		antColony = append(antColony, response)
	}
	garden.antColony = antColony
	return  garden
}

func (garden Garden) antsAddPheromone() (Garden) {
	for _,ant := range garden.antColony {
		garden = ant.AddPheromone(garden)
	}
	return garden
}

func (garden Garden) GetBestRoute() Ant{
	var bestRoute float64
	var bestAnt   Ant
	bestRoute = garden.antColony[0].GetRouteDistante()
	for _, ant := range garden.antColony {
		if ant.GetRouteDistante() < bestRoute {
			bestRoute = ant.GetRouteDistante()
			bestAnt = ant
		}
	}
	return bestAnt
}