package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"test_ecapture_backend/api/rest/handler/sessions_handler"
	"test_ecapture_backend/api/rest/handler/usuarios_handler"

	"test_ecapture_backend/api/rest/handler/perfiles_handler"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func routes(db *sqlx.DB, loggerHttp bool, allowedOrigins string) *fiber.App {
	app := fiber.New()

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST",
	}))
	if loggerHttp {
		app.Use(logger.New())
	}
	TxID := uuid.New().String()

	perfiles_handler.PerfilesRouter(app, db, TxID)
	usuarios_handler.UsuariosRouter(app, db, TxID)
	sessions_handler.SessionRouter(app, db, TxID)
	return app
}
