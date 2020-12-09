package service

var graphSingleton *Graph = nil

func (graph *Graph) getCity(cityName string) *City {
	for _, currentCity := range graph.Cities {
		if currentCity.Name == cityName {
			return currentCity
		}
	}
	return nil
}

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
	isPreset := false
	originCity := graph.getCity(origin.Name)
	destinationCity := graph.getCity(destination.Name)

	if originCity == nil {
		graph.AddCity(origin)
		originCity = origin
	}

	if destinationCity == nil {
		graph.AddCity(destination)
		destinationCity = destination
	}

	for _, currentRoad := range originCity.Connections {
		if currentRoad.Target == destinationCity {
			isPreset = true
			break
		}
	}
	if !isPreset {
		newRoad := &Road{Target: destinationCity, Cost: cost}
		originCity.Connections = append(originCity.Connections, newRoad)
	}
}
