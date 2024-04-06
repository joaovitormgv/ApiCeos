package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ApiCeos/app/handlers"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(app *fiber.App, h *handlers.Handlers) {
	// Rota para criar um novo usuário
	app.Post("users", h.CreateUserHandler)

	app.Get("users/:id", h.GetUserHandler)

	app.Get("users", h.GetUsersHandler)

	app.Put("users/:id", h.UpdateUserHandler)

	app.Delete("users/:id", h.DeleteUserHandler)
}
