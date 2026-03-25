package main

import (
	"os"

	"el-maistro/internal/routes"
	"el-maistro/pkg/db"

	"github.com/gofiber/fiber/v3"
	"github.com/subosito/gotenv"
)

func main() {
	// loading enviroment variables
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}

	database, err := db.NewSQLiteConnection("elmaistro.db")
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(database); err != nil {
		panic("could not migrate database: " + err.Error())
	}

	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	app.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.InitCustomersEndpoints(app, database)

	app.Listen(port)
}
