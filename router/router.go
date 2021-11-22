package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vikas/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.GetAllProducts)
	app.Get("/:id", controllers.GetProduct)
	app.Post("/", controllers.CreateProduct)
	app.Delete("/:id", controllers.DeleteProduct)
}
