package main

import (
	"log"
	"os"
	"run_island/apps/helper"
	"run_island/apps/models"
	"run_island/apps/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//database
	helper.ConnectDB()
	helper.DB.AutoMigrate(
		models.Kategori{},
	)

	//engine
	engine := html.New("./views", ".html")
	engine.AddFunc("number_format", helper.NumberFormat)
	r := fiber.New(fiber.Config{
		Views: engine,
	})

	//routes
	r.Static("public", "./public")
	routes.MainRoute(r)
	routes.KategoriR(r)

	r.Listen(os.Getenv("app_port"))
}
