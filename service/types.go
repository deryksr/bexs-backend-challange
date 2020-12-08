package service

type Graph struct {
	Cities []*City
}

type Road struct {
	Soruce *City
	Target *City
	Cost   int
}

type City struct {
	Name        string
	Visited     bool
	Connections []*Road
}
