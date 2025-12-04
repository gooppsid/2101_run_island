package routes

import (
	"run_island/apps/controllers"

	"github.com/gofiber/fiber/v2"
)

func MainRoute(r *fiber.App) {
	r.Get("/", controllers.Index)
	r.Get("registrasi/:slug", controllers.Registrasi)
	r.Get("bayarDulu/:phone", controllers.BayarDulu)
	r.Get("tiketku/:uniqid", controllers.Tiketku)
}
