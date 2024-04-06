package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaovitormgv/ApiCeos/app/handlers"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(app *fiber.App, h *handlers.Handlers) {
	// Rota para criar um novo usuário
	app.Post("users", h.CreateUserHandler)

	// Rota para encontrar um usuário por ID
	app.Get("users/:id", h.GetUserHandler)

	// Rota para encontrar todos os usuários
	app.Get("users", h.GetUsersHandler)

	// Rota para atualizar um usuário por id
	app.Put("users/:id", h.UpdateUserHandler)

	// Rota para deletar um usuário por id
	app.Delete("users/:id", h.DeleteUserHandler)
}
