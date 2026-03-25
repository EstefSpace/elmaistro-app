package routes

import (
	"database/sql"

	"el-maistro/internal/repository"

	"github.com/gofiber/fiber/v3"
)

func InitCustomersEndpoints(app *fiber.App, db *sql.DB) {
	_ = repository.NewCustomerRepository(db)

	customersGroup := app.Group("/api/customers")
	customersGroup.Get("/", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})
}
