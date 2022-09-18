package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zazhedho/gorental/src/routers"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err)
	}

	APP_PORT := os.Getenv("APP_PORT")
	fmt.Println("App running on port", APP_PORT)
	http.ListenAndServe(APP_PORT, mainRoute)

}
