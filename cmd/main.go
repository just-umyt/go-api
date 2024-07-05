package main

import (
	"log"

	"github.com/just-umyt/go-api/internal/router"
)

func main() {
	app := router.NewApp()

	log.Fatal(app.Listen(":8080"))
}
