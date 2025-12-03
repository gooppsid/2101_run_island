package routes

import (
	"run_island/apps/controllers"

	"github.com/gofiber/fiber/v2"
)

func MainRoute(r *fiber.App) {
	r.Get("/", controllers.Index)
}
