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

func (graph *Graph) GetCity(cityName string) *City {
	return nil
}

func (graph *Graph) AddCity(city *City) {
	graph.Cities = append(graph.Cities, city)
}

func (graph *Graph) AddRoad(origin, destination *City, cost int) {
}
