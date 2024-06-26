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
	err := user.UnmarshalData(c.Body())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Falha em analisar o corpo da requisição",
		})
	}
	// Validar os dados do usuário
	// Verificar se o email do usuário é vazio
	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O email é obrigatório",
		})
	}

	// Verificar se o email do usuário já existe
	row := h.DB.QueryRow("SELECT id FROM users WHERE email = $1", user.Email)
	var id int
	err = row.Scan(&id)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email já cadastrado",
		})
	}

	// Verificar se o primeiro nome do usuário é vazio
	if user.FirstName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O primeiro nome é obrigatório",
		})
	}

	// Criar um novo usuário no banco de dados
	row = h.DB.QueryRow("INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id", user.FirstName, user.LastName, user.Email)

	// Obter o ID do novo usuário
	err = row.Scan(&id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Retornar mensagem de sucesso e dados do usuário criado
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Usuário criado com sucesso",
		"id":         id,
		"first_name": user.FirstName,
		"email":      user.Email,
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

	// Converter o usuário para JSON usando MarshalData
	userJSON, err := user.MarshalData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao converter usuário para JSON",
		})
	}

	// Retornar dados do usuário encontrado
	return c.Status(fiber.StatusOK).Send(userJSON)
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

func (h *Handlers) UpdateUserHandler(c *fiber.Ctx) error {
	// Obter o ID do usuário
	id := c.Params("id")

	// Encontrar o usuário no banco de dados
	row := h.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)

	// Criar um novo struct usuário
	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Usuário não encontrado",
		})
	}

	// Atualizar os dados do usuário
	err = user.UnmarshalData(c.Body())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Falha em analisar o corpo da requisição",
		})
	}

	// Validar os dados do usuário
	// Verificar se o email do usuário é vazio
	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O email é obrigatório",
		})
	}

	// Verificar se o email do usuário já existe
	row = h.DB.QueryRow("SELECT id FROM users WHERE email = $1 AND id != $2", user.Email, id)
	var existingID int
	err = row.Scan(&existingID)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email já cadastrado",
		})
	}

	// Verificar se o primeiro nome do usuário é vazio
	if user.FirstName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O primeiro nome é obrigatório",
		})
	}

	// Atualizar o usuário no banco de dados
	_, err = h.DB.Exec("UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4", user.FirstName, user.LastName, user.Email, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao atualizar usuário",
		})
	}

	// Retornar mensagem de sucesso e dados do usuário atualizado
	return c.JSON(fiber.Map{
		"message": "Usuário atualizado com sucesso",
		"user":    user,
	})
}

func (h *Handlers) DeleteUserHandler(c *fiber.Ctx) error {
	// Obter o ID do usuário
	id := c.Params("id")

	// Excluir o usuário do banco de dados
	_, err := h.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Falha ao excluir usuário",
		})
	}

	// Retornar mensagem de sucesso
	return c.JSON(fiber.Map{
		"message": "Usuário excluído com sucesso",
	})
}
