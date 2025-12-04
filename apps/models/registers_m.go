package models

import "gorm.io/gorm"

type Registers struct {
	gorm.Model
	Uniqid   string
	Funrun   string
	Nama     string
	Email    string
	Phone    string
	Ktp      string
	Usia     string
	Goldar   string
	Nama1    string
	Phone1   int
	Alamat   string
	Penyakit string
	Harga    int
	Status   string
}
