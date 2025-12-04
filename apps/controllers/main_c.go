package controllers

import (
	"run_island/apps/helper"
	"run_island/apps/models"

	"github.com/gofiber/fiber/v2"
)

// index
func Index(c *fiber.Ctx) error {
	var kategori []models.Kategori
	helper.DB.Where("status = ?", "Show").Find(&kategori)

	return c.Render("main/index", fiber.Map{
		"title":    "",
		"kategori": kategori,
	}, "main/layout")
}

// registrasi
func Registrasi(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var kategori models.Kategori
	helper.DB.Where("slug = ?", slug).First(&kategori)

	usia := make([]int, 0)
	for i := 17; i <= 55; i++ {
		usia = append(usia, i)
	}

	if kategori.Slug == "" {
		return c.SendString("404 HALAMAN TIDAK ADA")
	}

	return c.Render("main/registrasi", fiber.Map{
		"title":    "| Registrasi",
		"kategori": kategori,
		"usia":     usia,
	}, "main/layout")
}

// bayar
func BayarDulu(c *fiber.Ctx) error {
	phone := c.Params("phone")

	var register models.Registers
	helper.DB.Where("phone = ?", phone).First(&register)

	return c.Render("main/bayar", fiber.Map{
		"title":    "| Bayar Tiketku",
		"register": register,
	}, "main/layout")
}

// tiketku
func Tiketku(c *fiber.Ctx) error {
	uniqid := c.Params("uniqid")

	var register models.Registers
	helper.DB.Where("uniqid = ?", uniqid).First(&register)

	if register.Status == "Pending" {
		helper.DB.Model(&models.Registers{}).Where("uniqid = ?", uniqid).Update(
			"status", "Paid",
		)
	}

	return c.Render("main/tiketku", fiber.Map{
		"title":    "| Tiketku",
		"register": register,
	}, "main/layout")
}
