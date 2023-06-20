package controllers

import (
	"context"
	"log"
	"os"

	"github.com/RianIhsan/goBioskop/database"
	"github.com/RianIhsan/goBioskop/models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func GetAll(c *fiber.Ctx) error {
  var films []models.Film

  database.DB.Find(&films)

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "data": films,
  }) 
}

func GetById(c *fiber.Ctx) error {
  var film models.Film

	filmId := c.Params("id")
	if filmId == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak boleh kosong",
		})
	}

	if err := database.DB.Where("id = ? ", filmId).First(&film).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data film tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data film ditemukan",
		"data":    film,
	})
}


func Add(c *fiber.Ctx) error {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Failed load env file")
  }

  var filmReq models.FilmReq

  _ = c.BodyParser(&filmReq)

  fileHeader, _ := c.FormFile("image")
  file, _ :=fileHeader.Open()

  log.Println(fileHeader.Filename)

  ctx := context.Background()

  urlCloudinary := os.Getenv("URLCLOUDINARY")
  cldService, _ := cloudinary.NewFromURL(urlCloudinary)

  res, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
  log.Println(res.SecureURL)

  var film models.Film
  film.Title = filmReq.Title
  film.Image = res.SecureURL
  film.Genre = filmReq.Genre
  film.Description = filmReq.Description
  film.Duration = filmReq.Duration
  film.Creator = filmReq.Creator
  film.ReleaseDate = filmReq.ReleaseDate

  if err := database.DB.Create(&film).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message":"Internal Server Error Create New Data",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "data": film,
  })
}

func Update(c *fiber.Ctx) error {
  if err := godotenv.Load(); err != nil {
    log.Fatal("Failed load env file")
  }

  var filmReq models.FilmReq

  _ = c.BodyParser(&filmReq)

  filmID := c.Params("id") // Mendapatkan parameter ID film dari URL

  var film models.Film

  // Cek apakah data film dengan ID tersebut ada dalam database
  if err := database.DB.First(&film, filmID).Error; err != nil {
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
      "message": "Film not found",
    })
  }

  fileHeader, _ := c.FormFile("image")
  if fileHeader != nil {
    file, _ := fileHeader.Open()

    ctx := context.Background()

    urlCloudinary := os.Getenv("URLCLOUDINARY")
    cldService, _ := cloudinary.NewFromURL(urlCloudinary)

    res, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
    film.Image = res.SecureURL
  }

  film.Title = filmReq.Title
  film.Genre = filmReq.Genre
  film.Description = filmReq.Description
  film.Duration = filmReq.Duration
  film.Creator = filmReq.Creator
  film.ReleaseDate = filmReq.ReleaseDate

  if err := database.DB.Save(&film).Error; err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
      "message": "Internal Server Error Update Data",
    })
  }

  return c.Status(fiber.StatusOK).JSON(fiber.Map{
    "data": film,
  })
}

func Delete(c *fiber.Ctx) error {
	film := models.Film{}

	filmId := c.Params("id")

	if err := database.DB.First(&film, filmId).Delete(&film).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Film gagal dihapus",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Film berhasil dihapus",
	})
}

