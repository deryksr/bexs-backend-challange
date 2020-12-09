package service

type Graph struct {
	Cities []*City
}

type Road struct {
	Target *City
	Cost   int
}

type City struct {
	Name        string
	Visited     bool
	Connections []*Road
}

type Route struct {
	Paths []string
	Cost  int
}

type RouteList []Route

func (p RouteList) Len() int           { return len(p) }
func (p RouteList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p RouteList) Less(i, j int) bool { return p[i].Cost < p[j].Cost }
