package api

import (
	"net/http"
)

func BestRoute(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		response.Write([]byte("It works!"))
		return
	default:
		return
	}
}
