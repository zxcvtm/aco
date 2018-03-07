package schemas

import (
	"math/rand"
	"time"
	"github.com/kr/pretty"
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
func (garden Garden) InitGarden(graph [][]Edge, alpha, beta float64) Garden {
	seed:= int64(time.Now().Nanosecond())
	source := rand.NewSource(seed)
	random := rand.New(source)

	garden.alpha = alpha
	garden.beta = beta
	garden.p = random.Float64()
	garden.graph = graph

	for x, edges := range garden.graph {
		for y, edge := range edges {
			edge.Pheromone = 1.0
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
	pretty.Println(garden)
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
func (garden Garden) antsAddPheromone() (Garden) {
	for _,ant := range garden.antColony {
		garden = ant.AddPheromone(garden)
	}
	return garden
}

func (garden Garden) GetBestRoute() float64{
	var bestRoute float64
	bestRoute = garden.antColony[0].GetRouteDistante()
	for _, ant := range garden.antColony {
		if ant.GetRouteDistante() < bestRoute {
			bestRoute = ant.GetRouteDistante()
		}
	}
	return bestRoute
}