package barang_my

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/repository/barang_repository"

	"gorm.io/gorm"
)

type barangMY struct {
	db *gorm.DB
}

func NewBarangMy(db *gorm.DB) barang_repository.BarangRepository {
	return &barangMY{
		db: db,
	}
}

func (b *barangMY) GetAllBarang() ([]*entity.Barang, custerrs.MessageErr) {
	var barang []*entity.Barang
	err := b.db.Find(&barang).Error

	if err != nil {
		return nil, custerrs.NewInternalServerError("error while getting barangs data")
	}

	return barang, nil
}

func (b *barangMY) GetBarangById(barangId int) (*entity.Barang, custerrs.MessageErr) {
	var barang entity.Barang
	result := b.db.First(&barang, barangId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, custerrs.NewNotFoundError("barang not found")
		}
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	return &barang, nil
}

func (b *barangMY) CreateBarang(barangPayload *entity.Barang) (*entity.Barang, custerrs.MessageErr) {
	result := b.db.Create(barangPayload)

	if result.Error != nil {
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	row := result.Row()

	var barang entity.Barang
	row.Scan(row, &barang)

	return &barang, nil
}

func (b *barangMY) UpdateBarangById(barangPayload *entity.Barang) custerrs.MessageErr {
	err := b.db.Model(barangPayload).Updates(entity.Barang{Barang: barangPayload.Barang, HargaBeli: barangPayload.HargaBeli, HargaJual: barangPayload.HargaJual}).Error

	if err != nil {
		return custerrs.NewInternalServerError("error while updating barang")
	}

	return nil
}

func (b *barangMY) DeleteBarangById(barangId int) custerrs.MessageErr {
	var barang entity.Barang
	result := b.db.First(&barang, barangId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return custerrs.NewNotFoundError("barang not found")
		}
		return custerrs.NewInternalServerError("something went wrong")
	}

	err := b.db.Delete(&barang).Error

	if err != nil {
		return custerrs.NewInternalServerError("something went wrong")
	}

	return nil
}
