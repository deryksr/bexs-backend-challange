package service

import (
	"errors"
	"sort"
)

func getRoutes(source, target *City, path []*City, allCost int, allPaths *RouteList) {
	path = append(path, source)

	if source == target {
		cities := parseCitiesToString(path)
		*allPaths = append(*allPaths, Route{Paths: []string{cities}, Cost: allCost})

	} else {
		source.Visited = true
		for _, road := range source.Connections {
			if !road.Target.Visited {
				currentCost := allCost + road.Cost
				getRoutes(road.Target, target, path, currentCost, allPaths)
			}
		}
		source.Visited = false
	}
}

func GetBestRoute(source, target string) Route {
	var bestRoute Route
	return bestRoute
}

func GetAllRoutes(source, target string) (RouteList, error) {
	allPaths := RouteList{}

	sourceCity := GetGraphInstance().getCity(source)
	targetCity := GetGraphInstance().getCity(target)

	if sourceCity == nil {
		return allPaths, errors.New("source <" + source + "> has not found")
	}

	if targetCity == nil {
		return allPaths, errors.New("target <" + target + "> has not found")
	}

	getRoutes(sourceCity, targetCity, make([]*City, 0), 0, &allPaths)

	if len(allPaths) < 1 {
		return allPaths, errors.New("None route has been found between " + source + " - " + target)
	}

	sort.Sort(allPaths)
	return allPaths, nil
}
