package api

import "net/http"

func StartServer(port string) {
	http.HandleFunc("/route/best", BestRoute)
	http.HandleFunc("/route/all", AllRoutes)
	http.HandleFunc("/route", AddRoute)

	http.ListenAndServe(":"+port, nil)
}
