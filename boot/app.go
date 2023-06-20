package boot

import (
	"log"
	"os"

	"github.com/RianIhsan/goBioskop/config"
	"github.com/RianIhsan/goBioskop/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)


func BootApp() {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Error access .env file")
  }

  port := os.Getenv("PORT")

  config.ConnectDB()
  config.RunMigrate()

  app := fiber.New()

  app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: false,
	}))

  routes.InitRoute(app)
  app.Listen(":"+port)

  
}
