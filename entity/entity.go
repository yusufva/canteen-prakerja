package entity

import (
	"canteen-prakerja/dto"
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	ID        int    `gorm:"primaryKey;not null" json:"id"`
	Barang    string `gorm:"unique;not null;varchar(255)" json:"barang"`
	HargaBeli int    `gorm:"not null" json:"harga_beli"`
	HargaJual int    `gorm:"not null" json:"harga_jual"`
	// ItemTerjual []Item         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"item terjual"`
	CreatedAt time.Time      `gorm:"default:current_timestamp(3)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Transaksi struct {
	ID         int            `gorm:"primaryKey;not null" json:"id"`
	Tanggal    time.Time      `gorm:"unique;default:current_timestamp(3)" json:"tanggal"`
	Items      []Item         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"item"`
	TotalHarga int            `gorm:"" json:"total"`
	CreatedAt  time.Time      `gorm:"default:current_timestamp(3)" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:current_timestamp(3)" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type Item struct {
	ID          int       `gorm:"primaryKey;not null" json:"id"`
	TransaksiID int       `gorm:"not null" json:"id_transaksi"`
	Barang      string    `gorm:"not null" json:"barang"`
	Jumlah      int       `gorm:"not null" json:"jumlah"`
	HargaJual   int       `gorm:"not null" json:"harga_jual"`
	HargaBeli   int       `gorm:"not null" json:"harga_beli"`
	Laba        int       `gorm:"not null" json:"laba"`
	TotalHarga  int       `gorm:"not null" json:"total_harga"`
	CreatedAt   time.Time `gorm:"default:current_timestamp(3)" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp(3)" json:"updated_at"`
	// DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type ItemRevenue struct {
	Barang       string `json:"barang"`
	JumlahTot    int    `json:"total_terjual"`
	HargaBeliTot int    `json:"total_modal"`
	HargaJualTot int    `json:"total_harga_jual"`
	Laba         int    `json:"laba"`
}

type TotalItemRevenue struct {
	JumlahSum    int `json:"total_terjual"`
	HargaBeliSum int `json:"total_modal"`
	HargaJualSum int `json:"total_harga_jual"`
	LabaSum      int `json:"laba"`
}

func (i *Item) EntityToItemResponseDto() dto.ItemResponse {
	return dto.ItemResponse{
		ID:          i.ID,
		TransaksiID: i.TransaksiID,
		Barang:      i.Barang,
		Jumlah:      i.Jumlah,
		HargaJual:   i.HargaJual,
		HargaBeli:   i.HargaBeli,
		Laba:        i.Laba,
		TotalHarga:  i.TotalHarga,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}
}

func (b *Barang) EntityToBarangResponseDto() dto.BarangResponse {
	return dto.BarangResponse{
		ID:        b.ID,
		Barang:    b.Barang,
		HargaBeli: b.HargaBeli,
		HargaJual: b.HargaJual,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func (t *Transaksi) EntityToTransaksiResponseDTO() dto.TransaksiResponse {
	return dto.TransaksiResponse{
		ID:         t.ID,
		Tanggal:    t.Tanggal,
		Items:      []dto.ItemResponse{},
		TotalHarga: t.TotalHarga,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}

func (ir *ItemRevenue) EntityToReportResponseDTO() dto.ReportResponse {
	return dto.ReportResponse{
		Barang:         ir.Barang,
		TotalTerjual:   ir.JumlahTot,
		TotalModal:     ir.HargaBeliTot,
		TotalHargaJual: ir.HargaJualTot,
		Laba:           ir.Laba,
	}
}

func (tir *TotalItemRevenue) EntityToTotalSumResponseDTO() dto.TotalSumResponse {
	return dto.TotalSumResponse{
		SumTerjual:   tir.JumlahSum,
		SumModal:     tir.HargaBeliSum,
		SumHargaJual: tir.HargaJualSum,
		SumLaba:      tir.LabaSum,
	}
}
