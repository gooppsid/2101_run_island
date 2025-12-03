package models

type Kategori struct {
	ID     int    `gorm:"primaryKey"`
	Uniqid string `gorm:"not null"`
	Nama   string `gorm:"not null"`
	Slug   string `gorm:"not null"`
	Harga  int    `gorm:"not null"`
	Limit  int    `gorm:"not null"`
	Status string `gorm:"not null"`
	Noted  string
}
