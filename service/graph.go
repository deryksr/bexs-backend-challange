package service

var graphSingleton *Graph = nil

func GetGraphSingleton() *Graph {
	if graphSingleton == nil {
		graphSingleton = new(Graph)
	}
	return graphSingleton
}

func CleanGraph() {
	graphSingleton = nil
}

func (graph *Graph) AddCity(city *City) {
	isPresent := false
	for _, currentCity := range graph.Cities {
		if currentCity.Name == city.Name {
			isPresent = true
			break
		}
	}
	if !isPresent {
		graph.Cities = append(graph.Cities, city)
	}
}

func (graph *Graph) AddRoad(origin, destination *City, cost int) {
}
