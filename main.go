package main

import (
	"banxs-backend-challange/api"
	"banxs-backend-challange/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	portServer := "3000"
	if len(os.Args) < 2 {
		fmt.Println("Error: none file has been received | go run main.go <file.csv>")
		return
	}

	csvFile, err := service.ReadCsvFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range csvFile {
		source := service.City{
			Name:        line[0],
			Visited:     false,
			Connections: nil,
		}
		target := service.City{
			Name:        line[1],
			Visited:     false,
			Connections: nil,
		}
		cost, _ := strconv.Atoi(line[2])

		service.GetGraphInstance().AddCity(&source)
		service.GetGraphInstance().AddCity(&target)
		service.GetGraphInstance().AddRoad(&source, &target, cost)
	}

	fmt.Println("App started at localhost:" + portServer)
	go api.StartServer(portServer)

	for {
		fmt.Print("please enter the route: ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		cities := strings.Split(text[:len(text)-1], "-")
		if len(cities) != 2 {
			fmt.Println("Invalid input string format | ex: <origin>-<destination>")
			continue
		}

		bestRoute, err := service.GetBestRoute(cities[0], cities[1])
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}

		if len(bestRoute.Paths) > 1 {
			fmt.Printf("%d best routes have been found\n", len(bestRoute.Paths))
			for _, currentPath := range bestRoute.Paths {
				fmt.Printf("route: %s > $%d\n", currentPath, bestRoute.Cost)
			}
		} else {
			fmt.Printf("best route: %s > $%d\n", bestRoute.Paths, bestRoute.Cost)
		}
	}

}
