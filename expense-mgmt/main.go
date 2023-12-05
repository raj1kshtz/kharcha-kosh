package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/raj1kshtz/kharcha-core/conf"
	"github.com/raj1kshtz/kharcha-kosh/expense-mgmt/constants"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("We are getting the env values")
	}

	conf.LoadConfiguration()

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		port = constants.DEFAULT_HTTP_PORT
	}

	ctx := context.Background()

}
