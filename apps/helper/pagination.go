package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Mengambil query parameter dari context Fiber
		page, _ := strconv.Atoi(c.Query("page", "1")) // Default page = 1
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size", "50")) // Default page_size = 10
		// Memastikan page_size tidak lebih dari 50 atau kurang dari 1
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 50 // Menetapkan default ke 10 jika tidak valid
		}

		// Hitung offset untuk pagination
		offset := (page - 1) * pageSize

		// Menggunakan offset dan limit untuk mengambil data
		return db.Offset(offset).Limit(pageSize)
	}
}
