package controllers

import (
	"run_island/apps/helper"
	"run_island/apps/models"

	"github.com/gofiber/fiber/v2"
)

// index
func Index(c *fiber.Ctx) error {
	var kategori []models.Kategori
	helper.DB.Where("status=?", "Show").Find(&kategori)

	return c.Render("main/index", fiber.Map{
		"title":    "",
		"kategori": kategori,
	}, "main/layout")
}
