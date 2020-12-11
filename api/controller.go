package api

import (
	"banxs-backend-challange/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func BestRoute(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		query := request.URL.Query()
		source := query.Get("source")
		target := query.Get("target")

		bestRoute, err := service.GetBestRoute(source, target)

		if err != nil {
			payload := struct{ Message string }{err.Error()}
			json, _ := json.Marshal(payload)
			response.Write(json)
		} else {
			response.Header().Set("Content-Type", "application/json")
			json, _ := json.Marshal(bestRoute)
			response.Write(json)
		}
		return
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("This method is not allowed"))
		return
	}
}

func AllRoutes(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		query := request.URL.Query()
		source := query.Get("source")
		target := query.Get("target")

		allRoutes, err := service.GetAllRoutes(source, target)

		if err != nil {
			payload := struct{ Message string }{err.Error()}
			json, _ := json.Marshal(payload)
			response.Write(json)
		} else {
			response.Header().Set("Content-Type", "application/json")
			json, _ := json.Marshal(allRoutes)
			response.Write(json)
		}
		return
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("This method is not allowed"))
		return
	}
}

func AddRoute(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		type NewRoad struct {
			Source string `json:"source"`
			Target string `json:"target"`
			Cost   int    `json:"cost"`
		}

		var received NewRoad
		body, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()

		if err != nil {
			http.Error(response, "Invalid payload", 400)
			return
		}

		err = json.Unmarshal(body, &received)

		if err != nil {
			http.Error(response, "Something went wrong!", 500)
			return
		}

		hasAdd := service.GetGraphInstance().AddRoad(
			&service.City{Name: received.Source, Visited: false, Connections: nil},
			&service.City{Name: received.Target, Visited: false, Connections: nil},
			received.Cost,
		)

		if hasAdd != nil {
			response.WriteHeader(http.StatusCreated)
			response.Write([]byte("The path has been added with success!"))
			service.WriteCsvFile(
				service.GetCsvFileName(),
				[]string{
					received.Source,
					received.Target,
					strconv.Itoa(received.Cost),
				},
			)
		} else {
			http.Error(response, "The path between "+received.Source+"-"+received.Target+" already exist!", 400)
		}
		return
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("This method is not allowed"))
		return
	}
}
