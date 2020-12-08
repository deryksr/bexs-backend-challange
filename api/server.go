package api

import "net/http"

func StartServer(port string) {
	http.HandleFunc("/best-route", BestRoute)

	http.ListenAndServe(":"+port, nil)
}
