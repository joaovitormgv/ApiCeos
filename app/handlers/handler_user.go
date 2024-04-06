package handlers

import (
	"database/sql"

	"github.com/joaovitormgv/ApiCeos/app/models"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	DB *sql.DB
}

// Função para lidar com requests http de criação de usuários
func (h *Handlers) CreateUserHandler(c *fiber.Ctx) error {
	user := &models.User{}
	user, err := models.UnmarshalData(c.Body())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Falha em analisar o corpo da requisição",
		})
	}

	// Criar um novo usuário no banco de dados
	row := h.DB.QueryRow("INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id", user.FirstName, user.LastName, user.Email)

	// Obter o ID do novo usuário
	var id int
	err = row.Scan(&id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Retornar mensagem de sucesso e dados do usuário criado
	return c.JSON(fiber.Map{
		"message": "Usuário criado com sucesso",
		"id":      id,
		"user":    user,
	})
}

// Função para lidar com requests http de encontrar usuário por ID
func (h *Handlers) GetUserHandler(c *fiber.Ctx) error {
	// Obter o ID do usuário
	id := c.Params("id")

	// Encontrar o usuário no banco de dados
	row := h.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)

	// Criar um novo usuário
	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Usuário não encontrado",
		})
	}

	// Retornar dados do usuário encontrado
	return c.JSON(user)
}

// Função para lidar com requests http de listar todos os usuários
func (h *Handlers) GetUsersHandler(c *fiber.Ctx) error {
	// Encontrar todos os usuários no banco de dados
	rows, err := h.DB.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao obter usuários",
		})
	}
	defer rows.Close()

	// Criar um slice de usuários
	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Falha ao obter usuários",
			})
		}
		users = append(users, user)
	}

	// Retornar dados dos usuários encontrados
	return c.JSON(users)
}
