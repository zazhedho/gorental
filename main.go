package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/zazhedho/gorental/src/config"
)

func main() {
	if err := config.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
