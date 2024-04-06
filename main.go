package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ApiCeos/app/handlers"
	"github.com/joaovitormgv/ApiCeos/app/routes"
)

func main() {
	// Obter string de conexão com o banco de dados da variável de ambiente
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "host=localhost port=5432 user=postgres dbname=ceos sslmode=disable"
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := &handlers.Handlers{DB: db}
	app := fiber.New()
	routes.SetupRoutes(app, h)
	app.Listen(":3000")
}
