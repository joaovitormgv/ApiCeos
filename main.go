package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ApiCeos/app/handlers"
	"github.com/joaovitormgv/ApiCeos/app/routes"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=ceos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := &handlers.Handlers{DB: db}
	app := fiber.New()
	routes.SetupRoutes(app, h)
	app.Listen(":3000")
}
