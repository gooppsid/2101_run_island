package routes

import (
	"run_island/apps/controllers"

	"github.com/gofiber/fiber/v2"
)

func KategoriR(r *fiber.App) {
	r.Get("admin/kategori", controllers.Kategori)
	r.Post("admin/simpan-kategori", controllers.SimpanKategori)
	r.Post("admin/update-kategori/:id", controllers.UpdateKategori)
	r.Post("admin/hapus-kategori/:id", controllers.HapusKategori)
}
