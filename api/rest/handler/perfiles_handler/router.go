package perfiles_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PerfilesRouter(app *fiber.App, db *sqlx.DB, tx string) {

	perfilHd := Handler{DB: db, TxID: tx}

	api := app.Group("/api")
	v1 := api.Group("/v1/perfiles")
	v1.Post("/", perfilHd.create)
	v1.Put("/:id", perfilHd.update)
	v1.Delete("/:id", perfilHd.delete)
	v1.Get("/:id", perfilHd.getByID)
	v1.Get("/", perfilHd.getAll)

}
