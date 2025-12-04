package controllers

import (
	"run_island/apps/helper"
	"run_island/apps/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

// kategori
func Kategori(c *fiber.Ctx) error {
	var kategori []models.Kategori
	helper.DB.Find(&kategori)

	return c.Render("admin/kategori", fiber.Map{
		"title":    "| Kategori",
		"kategori": kategori,
	}, "main/layout")
}

// simpan kategori
func SimpanKategori(c *fiber.Ctx) error {
	helper.DB.Model(&models.Kategori{}).Create(map[string]interface{}{
		"funrun": c.FormValue("funrun"),
		"nama":   c.FormValue("nama"),
		"slug":   slug.Make(c.FormValue("nama")),
		"harga":  c.FormValue("harga"),
		"limit":  c.FormValue("limit"),
		"status": c.FormValue("status"),
		"noted":  c.FormValue("noted"),
	})

	return c.Redirect("/admin/kategori")
}

// update kategori
func UpdateKategori(c *fiber.Ctx) error {
	id := c.Params("id")

	helper.DB.Model(&models.Kategori{}).Where("id=?", id).Updates(map[string]interface{}{
		"funrun": c.FormValue("funrun"),
		"nama":   c.FormValue("nama"),
		"slug":   slug.Make(c.FormValue("nama")),
		"harga":  c.FormValue("harga"),
		"limit":  c.FormValue("limit"),
		"status": c.FormValue("status"),
		"noted":  c.FormValue("noted"),
	})

	return c.Redirect("/admin/kategori")
}

// hapus kategori
func HapusKategori(c *fiber.Ctx) error {
	id := c.Params("id")

	helper.DB.Where("id=?", id).Unscoped().Delete(&models.Kategori{})
	return c.Redirect("/admin/kategori")
}
