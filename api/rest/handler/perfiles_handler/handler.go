package perfiles_handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"test_ecapture_backend/internal/logger"
	"test_ecapture_backend/pkg/perfiles"

	"net/http"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB   *sqlx.DB
	TxID string
}

func (h *Handler) create(c *fiber.Ctx) error {
	res := PerfilResponse{}
	m := PerfilRequest{}
	
	err := c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear perfil: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	servicePerfil := perfiles.NewPerfilService(perfiles.NewPerfilPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)

	err = servicePerfil.Create(m.Name)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo crear el perfil: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	res.Name = m.Name
	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) update(c *fiber.Ctx) error {
	res := PerfilResponse{}
	m := PerfilRequest{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	err = c.BodyParser(&m)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo leer el Modelo crear perfil: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	servicePerfil := perfiles.NewPerfilService(perfiles.NewPerfilPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = servicePerfil.Update(id, m.Name)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo actualizar el perfil: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	res.Name = m.Name
	return c.Status(http.StatusOK).JSON(res)

}

func (h *Handler) delete(c *fiber.Ctx) error {
	res := PerfilResponse{}
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	servicePerfil := perfiles.NewPerfilService(perfiles.NewPerfilPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	err = servicePerfil.Delete(id)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo eliminar el perfil: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(res)

}

func (h *Handler) getByID(c *fiber.Ctx) error {
	res := PerfilResponse{}

	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo convertir el id a int: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	servicePerfil := perfiles.NewPerfilService(perfiles.NewPerfilPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	p, err := servicePerfil.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusOK).JSON(res)
		}
		logger.Error.Printf(h.TxID, "no se pudo obtener el perfil: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(p)

}

func (h *Handler) getAll(c *fiber.Ctx) error {
	var res []PerfilResponse

	servicePerfil := perfiles.NewPerfilService(perfiles.NewPerfilPsqlRepository(h.DB, nil, h.TxID), nil, h.TxID)
	ps, err := servicePerfil.GetAll()
	if err != nil {
		logger.Error.Printf(h.TxID, "no se pudo obtener los perfiles: %v", err)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	return c.Status(http.StatusOK).JSON(ps)

}
