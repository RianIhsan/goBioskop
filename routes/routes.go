package routes

import (
	"github.com/RianIhsan/goBioskop/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
  app.Post("/api/films", controllers.Add)
  app.Get("/api/films", controllers.GetAll)
  app.Get("/api/films/:id", controllers.GetById)
  app.Put("/api/films/:id", controllers.Update)
  app.Delete("/api/films/:id", controllers.Delete)
}
