package usuarios_handler

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/pkg/usuarios"
)

type Handler struct {
	DB   *sqlx.DB
	TxID string
}

func (h *Handler) create(c *fiber.Ctx) error {
	res := UserResponse{}
	m := UserRequest{}

	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear usuario: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	serviceUser := usuarios.NewUserService(usuarios.NewUserPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)

	err = serviceUser.Create(m.Id_perfil, m.Usuario, m.Password, m.Nombres, m.Apellidos)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo crear el usuario: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	return c.Status(http.StatusCreated).JSON(UserResponse{
		Id_perfil: m.Id_perfil,
		Usuario:   m.Usuario,
		Password:  m.Password,
		Nombres:   m.Nombres,
		Apellidos: m.Apellidos,
	})
}

func (h *Handler) update(c *fiber.Ctx) error {
	res := UserResponse{}
	m := UserRequest{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	err = c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear usuario: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	fmt.Println(m)
	serviceUser := usuarios.NewUserService(usuarios.NewUserPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = serviceUser.Update(id, m.Id_perfil, m.Usuario, m.Password, m.Nombres, m.Apellidos)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo actualizar el usuario: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	return c.Status(http.StatusOK).JSON(UserResponse{
		Id_perfil: m.Id_perfil,
		Usuario:   m.Usuario,
		Password:  m.Password,
		Nombres:   m.Nombres,
		Apellidos: m.Apellidos,
	})

}

func (h *Handler) delete(c *fiber.Ctx) error {
	res := UserResponse{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	serviceUser := usuarios.NewUserService(usuarios.NewUserPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = serviceUser.Delete(id)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo eliminar el usuario: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(res)

}

func (h *Handler) getByID(c *fiber.Ctx) error {
	res := UserResponse{}

	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	serviceUser := usuarios.NewUserService(usuarios.NewUserPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	p, err := serviceUser.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusOK).JSON(res)
		}
		logger.Error.Printf(h.TxID, "no se pudo obtener el usuario: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(p)

}

func (h *Handler) getAll(c *fiber.Ctx) error {
	var res []UserResponse

	serviceUser := usuarios.NewUserService(usuarios.NewUserPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	ps, err := serviceUser.GetAll()
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo obtener los usuarios: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(ps)

}
