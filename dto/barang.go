package dto

import "time"

type NewBarangRequest struct {
	Barang    string `json:"barang" valid:"required~Barang tidak boleh kosong"`
	HargaBeli int    `json:"harga_beli" valid:"required~Harga Beli tidak boleh kosong"`
	HargaJual int    `json:"harga_jual" valid:"required~Harga Jual tidak boleh kosong"`
}

type NewBarangResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type BarangResponse struct {
	ID        int       `json:"id"`
	Barang    string    `json:"barang"`
	HargaBeli int       `json:"harga_beli"`
	HargaJual int       `json:"harga_jual"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBarangResponse struct {
	Result     string           `json:"result"`
	Message    string           `json:"message"`
	Statuscode int              `json:"statuscode"`
	Data       []BarangResponse `json:"data"`
}
