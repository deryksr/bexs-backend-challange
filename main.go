package main

import (
	"banxs-backend-challange/api"
	"fmt"
	"os"
)

func main() {
	portServer := "3000"
	if len(os.Args) < 2 {
		fmt.Println("Error: none file has been received | go run main.go <file.csv>")
		return
	}

	fmt.Println("App started at localhost:" + portServer)
	api.StartServer(portServer)

}
