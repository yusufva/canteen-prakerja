package transaksi_my

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/repository/transaksi_repository"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type transaksiMY struct {
	db *gorm.DB
}

func NewTransaksiMy(db *gorm.DB) transaksi_repository.TransaksiRepository {
	return &transaksiMY{
		db: db,
	}
}

func (t *transaksiMY) GetAllTransaksi() ([]*entity.Transaksi, custerrs.MessageErr) {
	var transaksi []*entity.Transaksi
	err := t.db.Preload("Items").Find(&transaksi).Error

	if err != nil {
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	return transaksi, nil
}

func (t *transaksiMY) GetTransaksiById(transaksiId int) (*entity.Transaksi, custerrs.MessageErr) {
	var transaksi entity.Transaksi
	result := t.db.Preload("Items").First(&transaksi, transaksiId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, custerrs.NewNotFoundError("transaksi not found")
		}
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	return &transaksi, nil
}

func (t *transaksiMY) GetTransaksiDateBetween(transaksiPayload transaksi_repository.DateBetween) ([]*entity.Transaksi, custerrs.MessageErr) {
	var transaksis []*entity.Transaksi

	cond := transaksi_repository.DateBetween{
		StartDate: fmt.Sprintf("%s 00:00:00", transaksiPayload.StartDate),
		EndDate:   fmt.Sprintf("%s 23:59:59", transaksiPayload.EndDate),
	}

	err := t.db.Where("tanggal BETWEEN ? AND ?", cond.StartDate, cond.EndDate).Preload("Items").Find(&transaksis).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, custerrs.NewNotFoundError("transaksi not found")
		}
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	return transaksis, nil
}

func (t *transaksiMY) CreateTransaksi(transaksiPayload *entity.Transaksi) (*entity.Transaksi, custerrs.MessageErr) {
	result := t.db.Clauses(clause.Returning{}).Create(transaksiPayload)

	if result.Error != nil {
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	// row := result.Row()

	// var transaction entity.Transaksi
	// row.Scan(row, &transaction)

	return transaksiPayload, nil
}

// func (t *transaksiMY) UpdateTransaksiById(transaksiId int, transaksiPayload *entity.Transaksi) custerrs.MessageErr {
// 	err := t.db.Model(transaksiPayload).Updates(entity.Transaksi{Tanggal: transaksiPayload.Tanggal, TotalHarga: transaksiPayload.TotalHarga}).Error

// 	if err != nil {
// 		custerrs.NewInternalServerError("something went wrong")
// 	}

// 	return nil
// }

func (t *transaksiMY) DeleteTransaksiById(transaksiId int) custerrs.MessageErr {
	var transaksi entity.Transaksi
	result := t.db.First(&transaksi, transaksiId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return custerrs.NewNotFoundError("transaksi not found")
		}
		return custerrs.NewInternalServerError("something went wrong")
	}

	err := t.db.Delete(&transaksi).Error

	if err != nil {
		return custerrs.NewInternalServerError("error while deleting transaction")
	}

	return nil
}
