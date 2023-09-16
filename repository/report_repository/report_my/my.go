package report_my

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/repository/report_repository"
	"fmt"

	"gorm.io/gorm"
)

type reportMy struct {
	db *gorm.DB
}

func NewReportMy(db *gorm.DB) report_repository.ReportRepository {
	return &reportMy{
		db: db,
	}
}

func (r *reportMy) GetReportDateBetween(reportPayload report_repository.DateBetween) ([]*entity.ItemRevenue, custerrs.MessageErr) {
	var items []*entity.ItemRevenue
	subquery1 := r.db.Model(&entity.Item{}).Select("items.barang, items.transaksi_id, items.jumlah, items.harga_beli, items.harga_jual, items.total_harga, items.laba")
	subquery2 := r.db.Model(&entity.Transaksi{}).Select("id, tanggal")

	cond := report_repository.DateBetween{
		StartDate: fmt.Sprintf("%s 00:00:00", reportPayload.StartDate),
		EndDate:   fmt.Sprintf("%s 23:59:59", reportPayload.EndDate),
	}

	err := r.db.Table("(?) A left join (?) B ON A.transaksi_id = B.id", subquery1, subquery2).Select("A.barang as barang, SUM(A.jumlah) as JumlahTot, SUM(A.jumlah*A.harga_beli) as HargaBeliTot, SUM(A.total_harga) as HargaJualTot, SUM(A.laba) as Laba").Where("B.tanggal BETWEEN ? AND ?", cond.StartDate, cond.EndDate).Group("A.barang").Find(&items).Error

	if err != nil {
		return nil, custerrs.NewInternalServerError("order went wrong")
	}

	return items, nil
}

func (r *reportMy) GetTotalReportDateBetween(reportPayload report_repository.DateBetween) ([]*entity.TotalItemRevenue, custerrs.MessageErr) {
	var total []*entity.TotalItemRevenue
	subquery := r.db.Model(&entity.Item{}).Select("transaksis.tanggal, items.transaksi_id, items.jumlah, items.harga_beli, items.total_harga, items.laba").Joins("left join transaksis ON items.transaksi_id = transaksis.id")

	cond := report_repository.DateBetween{
		StartDate: fmt.Sprintf("%s 00:00:00", reportPayload.StartDate),
		EndDate:   fmt.Sprintf("%s 23:59:59", reportPayload.EndDate),
	}

	err := r.db.Table("(?) A", subquery).Select("SUM(A.jumlah) as JumlahSum, SUM(A.jumlah*A.harga_beli) as HargaBeliSum, SUM(A.total_harga) as HargaJualSum, SUM(A.laba) as LabaSum").Where("A.tanggal BETWEEN ? AND ?", cond.StartDate, cond.EndDate).Find(&total).Error

	if err != nil {
		return nil, custerrs.NewInternalServerError("total went wrong")
	}

	// total[0].Barang = "total"

	return total, nil
}
