package main

import (
	"log"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/config"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitEnv()
	config.InitDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Cookie",
		AllowCredentials: true,
	}))

	app.Static("/uploads", "./uploads")
	routes.Setup(app)

	log.Fatal(app.Listen(config.Env.Host + ":" + config.Env.Port))
}
