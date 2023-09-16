package transaksi_repository

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
)

type DateBetween struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type TransaksiRepository interface {
	GetAllTransaksi() ([]*entity.Transaksi, custerrs.MessageErr)
	GetTransaksiById(transaksiId int) (*entity.Transaksi, custerrs.MessageErr)
	GetTransaksiDateBetween(transaksiPayload DateBetween) ([]*entity.Transaksi, custerrs.MessageErr)
	CreateTransaksi(transaksiPayload *entity.Transaksi) (*entity.Transaksi, custerrs.MessageErr)
	// UpdateTransaksiById(transaksiId int, transaksiPayload *entity.Transaksi) custerrs.MessageErr
	DeleteTransaksiById(transaksiId int) custerrs.MessageErr
}
