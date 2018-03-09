package schemas

import (
	"errors"
	"math"
	"time"
	"math/rand"
)

type (
	Ant struct {
		visited         []int
		route           map[int]float64
		pheromoneOnEdge []Edge
		q               float64
	}
)
func (ant Ant) InitAnt(startNode int, Q float64) Ant{
	distance := 0.0
	ant.visited = append(ant.visited, startNode)
	ant.route   = make(map[int]float64)
	ant.route[startNode] = distance
	ant.q = Q
	return ant;
}

func (ant Ant) getRoute(node int) (float64, error) {
	if _, ok := ant.route[node]; !ok {
		return ant.route[node], nil
	}
	return  0, errors.New("Route not found")
}
func (ant Ant) probability(garden Garden, to int) float64 {
	var(
		from                int = ant.visited[len(ant.visited)-1]
		probability         float64 = 0.0
		pheromoneTo         float64 = garden.graph[from][to].Pheromone
		visibilityTo        float64 = garden.Visibility(from,to)
		remainigRoute       float64 = 0.0
	)

	for index, edge := range garden.graph[from] {
		_, err := ant.getRoute(index)
		if err == nil {
			remainigPheromone := edge.Pheromone
			remainigVisibility := garden.Visibility(from, index)

			remainigRoute += math.Pow(remainigPheromone, garden.alpha)* math.Pow(remainigVisibility, garden.beta)
		}
	}
	probability = (math.Pow(pheromoneTo, garden.alpha) * math.Pow(visibilityTo, garden.beta))/remainigRoute
	return probability
}

func (ant Ant) Move(garden Garden) (Ant){
	var (
		startNode    int = ant.visited[len(ant.visited)-1]
		selectedNode int
	)
	selectedNode = ant.getNextNode(garden)
	ant = ant.moveToNextNode(selectedNode, garden.GetDistance(startNode, selectedNode))
	return ant;
}
func (ant Ant) MoveAsync(garden Garden) (<- chan Ant){
	c := make(chan Ant)
	go func() {
		var (
			startNode    int = ant.visited[len(ant.visited)-1]
			selectedNode int
		)
		selectedNode = ant.getNextNode(garden)
		ant = ant.moveToNextNode(selectedNode, garden.GetDistance(startNode, selectedNode))
		c <- ant;
	}()
	return c;
}

func (ant Ant) getNextNode(garden Garden) int {
	seed:= int64(time.Now().Nanosecond())
	source := rand.NewSource(seed)
	random := rand.New(source)
	selected := random.Float64()
	sum := 0.0
	for node, _ := range garden.graph  {
		_, err := ant.getRoute(node)
		if err == nil {
			sum += ant.probability(garden, node)
			if sum >= selected {
				return node
			}
		}
	}
	return 0
}

func (ant Ant) moveToNextNode(to int, distance float64) (Ant) {
	from:= ant.visited[len(ant.visited)-1]
	ant.visited = append(ant.visited, to)
	ant.route[to] = ant.route[from] + distance
	return ant
}

func (ant Ant) AddPheromone(garden Garden) (Garden) {
	var (
		lastNodeVisited int
		totalDistance   float64
		deltaPheromone  float64
	)
	lastNodeVisited = ant.visited[len(ant.visited)-1]
	totalDistance = ant.route[lastNodeVisited]
	deltaPheromone = ant.q/totalDistance

	for index, endNode := range ant.visited {
		if index != 0 {
			startNode := ant.visited[index-1]
			garden.graph[startNode][endNode].Pheromone = (1-garden.p) * garden.graph[startNode][endNode].Pheromone + deltaPheromone
		}
	}
	return garden
}

func (ant Ant) GetRouteDistante() float64 {
	lenght := len(ant.visited) -1
	if lenght < 0 {
		lenght = 0
	}
	lastNode := ant.visited[lenght]
	return ant.route[lastNode]
}

func (ant Ant) GetVisited() []int {
	return ant.visited
}