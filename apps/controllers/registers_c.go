package controllers

import (
	"run_island/apps/helper"
	"run_island/apps/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// simpan register
func SimpanRegister(c *fiber.Ctx) error {
	funrun := c.Params("funrun")
	harga := c.Params("harga")
	uniqid := helper.UniqID()
	phone := c.FormValue("code") + c.FormValue("phone")

	helper.DB.Model(&models.Registers{}).Create(map[string]interface{}{
		"uniqid":   uniqid,
		"funrun":   funrun,
		"nama":     c.FormValue("nama"),
		"email":    c.FormValue("email"),
		"phone":    phone,
		"ktp":      c.FormValue("ktp"),
		"usia":     c.FormValue("usia"),
		"goldar":   c.FormValue("goldar"),
		"nama1":    c.FormValue("nama1"),
		"phone1":   c.FormValue("phone1"),
		"alamat":   c.FormValue("alamat"),
		"penyakit": c.FormValue("penyakit"),
		"harga":    harga,
		"status":   "Pending",
	})

	return c.Redirect("/bayarDulu/" + phone)
}

// bayar
func Bayar(c *fiber.Ctx) error {
	uniqid := c.Params("uniqid")

	var register models.Registers
	helper.DB.Where("uniqid = ?", uniqid).First(&register)

	harga := strconv.Itoa(register.Harga)

	helper.Payment(c, register.Phone, "Fun Run "+register.Funrun+" Nama: "+register.Nama, harga, c.BaseURL(), uniqid)

	return nil
}
