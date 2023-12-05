package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/raj1kshtz/kharcha-core/conf"
	"github.com/raj1kshtz/kharcha-kosh/expense-mgmt/constants"
	"github.com/raj1kshtz/kharcha-kosh/expense-mgmt/routing"
	"log"
	"net/http"
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

	router := routing.NewRouter(ctx)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Failed to serve")
	}

}
