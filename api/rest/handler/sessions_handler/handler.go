package sessions_handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/pkg/usuarios"
)

type Handler struct {
	DB   *sqlx.DB
	TxID string
}

func (h *Handler) login(c *fiber.Ctx) error {
	res := LoginResponse{}
	m := LoginRequest{}

	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear usuario: %v", err)
		return c.Status(http.StatusForbidden).JSON(res)
	}

	serviceUser := usuarios.NewUserService(usuarios.NewUserPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	user, err := serviceUser.GetByUser(m.Usuario)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusForbidden).JSON(res)
		}
		logger.Error.Printf(h.TxID, "no se pudo obtener el usuario: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if m.Password != user.Password {
		return c.Status(http.StatusForbidden).JSON(res)
	}
	user.Password = ""
	return c.Status(http.StatusOK).JSON(user)

}
