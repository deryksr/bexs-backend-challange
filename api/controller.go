package api

import (
	"banxs-backend-challange/service"
	"encoding/json"
	"net/http"
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
