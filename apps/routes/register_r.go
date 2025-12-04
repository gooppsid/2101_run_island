package routes

import (
	"run_island/apps/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterR(r *fiber.App) {
	r.Post("simpanRegister/:funrun/:harga", controllers.SimpanRegister)
	r.Post("bayar/:uniqid", controllers.Bayar)
}
