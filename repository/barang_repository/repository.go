package barang_repository

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
)

type BarangRepository interface {
	GetAllBarang() ([]*entity.Barang, custerrs.MessageErr)
	GetBarangById(barangId int) (*entity.Barang, custerrs.MessageErr)
	CreateBarang(barangPayload *entity.Barang) (*entity.Barang, custerrs.MessageErr)
	UpdateBarangById(barangPayload *entity.Barang) custerrs.MessageErr
	DeleteBarangById(barangId int) custerrs.MessageErr
}
