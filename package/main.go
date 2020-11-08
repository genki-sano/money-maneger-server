package main

import (
	"os"

	"github.com/genki-sano/money-maneger-server/package/infrastructure"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := infrastructure.Router.Run(":" + port); err != nil {
		panic(err)
	}
}
