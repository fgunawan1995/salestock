package main
import (
	_ "github.com/mattn/go-sqlite3"
)
//Model
type Items struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	SKU string `gorm:"not null" form:"sku" json:"sku"`
	NamaItem  string `gorm:"not null" form:"nama_item" json:"nama_item"`
	JumlahSekarang  int `gorm:"not null" form:"jumlah_sekarang" json:"jumlah_sekarang"`
}
type ItemsIn struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Waktu string `gorm:"not null" form:"waktu" json:"waktu"`
	SKU string `gorm:"not null" form:"sku" json:"sku"`
	NamaBarang  string `gorm:"not null" form:"nama_barang" json:"nama_barang"`
	JumlahPemesanan  int `gorm:"not null" form:"jumlah_pemesanan" json:"jumlah_pemesanan"`
	JumlahDiterima  int `gorm:"not null" form:"jumlah_diterima" json:"jumlah_diterima"`
	HargaBeli  int `gorm:"not null" form:"harga_beli" json:"harga_beli"`
	Total  int `gorm:"not null" form:"total" json:"total"`
	NomorKwitansi string `gorm:"not null" form:"nomor_kwitansi" json:"nomor_kwitansi"`
	Catatan string `gorm:"not null" form:"catatan" json:"catatan"`
}
type ItemsOut struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Waktu string `gorm:"not null" form:"waktu" json:"waktu"`
	SKU string `gorm:"not null" form:"sku" json:"sku"`
	NamaBarang  string `gorm:"not null" form:"nama_barang" json:"nama_barang"`
	JumlahKeluar  int `gorm:"not null" form:"jumlah_keluar" json:"jumlah_keluar"`
	HargaJual  int `gorm:"not null" form:"harga_jual" json:"harga_jual"`
	Total  int `gorm:"not null" form:"total" json:"total"`
	Catatan string `gorm:"not null" form:"catatan" json:"catatan"`
}

type ItemsReport struct {
	SKU string `form:"sku" json:"sku"`
	NamaItem  string `form:"nama_item" json:"nama_item"`
	Jumlah int `form:"jumlah" json:"jumlah"`
	Rata int `form:"rata" json:"rata"`
	Total int `form:"total" json:"total"`
}

type SalesReport struct {
	IdPesanan string `form:"id_pesanan" json:"id_pesanan"`
	Waktu string `form:"waktu" json:"waktu"`
	SKU string `form:"sku" json:"sku"`
	NamaBarang  string `form:"nama_barang" json:"nama_barang"`
	Jumlah  int `form:"jumlah" json:"jumlah"`
	HargaJual  int `form:"harga_jual" json:"harga_jual"`
	Total int `form:"total" json:"total"`
	HargaBeli  int `form:"harga_beli" json:"harga_beli"`
	Laba int `form:"laba" json:"laba"`
}
