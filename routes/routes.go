package routes

import (
	"github.com/RianIhsan/goBioskop/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
  app.Post("/api/create", controllers.Add)
  app.Get("/api/reads", controllers.GetAll)
  app.Get("/api/read/:id", controllers.GetById)
  app.Put("/api/update/:id", controllers.Update)
  app.Delete("/api/delete/:id", controllers.Delete)
}
