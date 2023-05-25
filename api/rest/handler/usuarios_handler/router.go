package usuarios_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UsuariosRouter(app *fiber.App, db *sqlx.DB, tx string) {

	userHd := Handler{DB: db, TxID: tx}

	api := app.Group("/api")
	v1 := api.Group("/v1/usuarios")
	v1.Post("/", userHd.create)
	v1.Put("/:id", userHd.update)
	v1.Delete("/:id", userHd.delete)
	v1.Get("/:id", userHd.getByID)
	v1.Get("/", userHd.getAll)

}
