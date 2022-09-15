package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zazhedho/gorental/src/routers"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("App running on port 8080")
	http.ListenAndServe(":8080", mainRoute)

}
