package routes

import (
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/config"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/controller"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/middleware"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/repository"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/service"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func setupV1(app *fiber.App) {
	db := config.DB
	v1 := app.Group("/v1")

	// Repositories
	productRepository := repository.Product{DB: db}
	userRepository := repository.User{DB: db}

	// Middleware
	authenticator := middleware.Authenticator
	authorizer := middleware.Authorizer(userRepository)
	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.Env.Secret.AccessToken)},
	})

	// Services
	productService := service.Product{Repository: &productRepository}
	authService := service.Auth{UserRepository: userRepository}

	productController := controller.Product{Service: &productService}
	productGroup := v1.Group("products")
	{
		productGroup.Get("", productController.FindAll)
		productGroup.Get(":id", productController.FindOne)
		productGroup.Post("", jwtMiddleware, authenticator, authorizer, productController.Create)
		productGroup.Patch(":id", jwtMiddleware, authenticator, authorizer, productController.Update)
		productGroup.Delete(":id", jwtMiddleware, authenticator, authorizer, productController.Delete)
	}

	authController := controller.Auth{Service: authService}
	authGroup := v1.Group("auth")
	{
		authGroup.Post("/sign-up", authController.SignUp)
		authGroup.Post("/sign-in", authController.SignIn)
		authGroup.Delete("/sign-out", jwtMiddleware, authenticator, authController.SignOut)
		authGroup.Delete("/cookie", authController.RemoveCookie)
		authGroup.Post("/refresh-token", authController.RefreshToken)
		authGroup.Get("/profile", jwtMiddleware, authenticator, authController.GetProfile)
		authGroup.Patch("/profile", jwtMiddleware, authenticator, authController.UpdateProfile)
	}
}
