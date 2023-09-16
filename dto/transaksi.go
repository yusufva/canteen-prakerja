package dto

import "time"

type NewTransaksiRequest struct {
	Tanggal    time.Time        `json:"tanggal"`
	Items      []NewItemRequest `json:"items"`
	TotalHarga int              `json:"total_harga" valid:"required~Total Harga tidak boleh kosong"`
}

type LatestTransaksiResponse struct {
	ID      int       `json:"id"`
	Tanggal time.Time `json:"tanggal"`
	Total   int       `json:"total"`
}

type NewTransaksiResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type TransaksiDateBetweenRequest struct {
	StartDate string `json:"start_date" valid:"required~Start Date tidak boleh kosong"`
	EndDate   string `json:"end_date" valid:"required~End Date tidak boleh kosong"`
}

type TransaksiResponse struct {
	ID         int            `json:"id"`
	Tanggal    time.Time      `json:"tanggal"`
	Items      []ItemResponse `json:"items"`
	TotalHarga int            `json:"total_harga"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type GetTransaksiResponse struct {
	Result     string              `json:"result"`
	Message    string              `json:"message"`
	StatusCode int                 `json:"statuscode"`
	Data       []TransaksiResponse `json:"data"`
}

type NewItemRequest struct {
	Barang     string `json:"barang"`
	Jumlah     int    `json:"jumlah"`
	HargaJual  int    `json:"harga_jual"`
	HargaBeli  int    `json:"harga_beli"`
	TotalHarga int    `json:"total_harga"`
}

type ItemResponse struct {
	ID          int       `json:"id_soldItems"`
	TransaksiID int       `json:"id_transaksi"`
	Barang      string    `json:"barang"`
	Jumlah      int       `json:"jumlah"`
	HargaJual   int       `json:"harga_jual"`
	HargaBeli   int       `json:"harga_beli"`
	Laba        int       `json:"laba"`
	TotalHarga  int       `json:"total_harga"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewItemResponse struct {
	Result     string         `json:"result"`
	Message    string         `json:"message"`
	StatusCode int            `json:"statuscode"`
	Data       []ItemResponse `json:"data"`
}
