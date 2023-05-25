package sessions_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SessionRouter(app *fiber.App, db *sqlx.DB, tx string) {

	sessionHd := Handler{DB: db, TxID: tx}

	api := app.Group("/api")
	v1 := api.Group("/v1/sessions")
	v1.Post("/", sessionHd.login)
}
