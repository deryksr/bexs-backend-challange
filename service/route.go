package service

import (
	"errors"
	"sort"
)

func parseCitiesToString(cities []*City) (result string) {
	for position, city := range cities {
		if position != len(cities)-1 {
			result += city.Name + " - "
		} else {
			result += city.Name
		}
	}
	return result
}

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

func GetBestRoute(source, target string) (Route, error) {
	bestRoute := Route{}
	allRoutes, err := GetAllRoutes(source, target)

	if err != nil {
		return bestRoute, err
	}

	bestRoute.Cost = allRoutes[0].Cost

	for _, currentRoute := range allRoutes {
		if bestRoute.Cost == currentRoute.Cost {
			bestRoute.Paths = append(bestRoute.Paths, currentRoute.Paths...)
		} else {
			break
		}
	}
	return bestRoute, nil
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
	pathsResult := RouteList{}

	for _, route := range allPaths {
		isPresent := false
		if len(pathsResult) != 0 {
			for index, _ := range pathsResult {
				if pathsResult[index].Cost == route.Cost {
					pathsResult[index].Paths = append(pathsResult[index].Paths, route.Paths...)
					isPresent = true
					break
				}
			}
		}
		if !isPresent {
			pathsResult = append(pathsResult, route)
		}
	}

	return pathsResult, nil
}
