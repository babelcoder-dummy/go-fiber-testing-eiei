package server

import (
	"log"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/config"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	App *fiber.App
}

func New() *Server {
	return &Server{App: fiber.New()}
}

func (s *Server) Setup() {
	config.InitEnv()
	config.InitDB()

	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Cookie",
		AllowCredentials: true,
	}))

	s.App.Static("/uploads", "./uploads")
	routes.Setup(s.App)
}

func (s *Server) Start() {
	log.Fatal(s.App.Listen(config.Env.Host + ":" + config.Env.Port))
}
