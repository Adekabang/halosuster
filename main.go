package main

import (
	"github.com/Adekabang/halosuster/pkg/configs"
	"github.com/Adekabang/halosuster/pkg/middleware"
	"github.com/Adekabang/halosuster/pkg/routes"
	"github.com/Adekabang/halosuster/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	// routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServer(app)
}
