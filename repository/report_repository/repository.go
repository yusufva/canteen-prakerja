package report_repository

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
)

type DateBetween struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ReportRepository interface {
	GetReportDateBetween(reportPayload DateBetween) ([]*entity.ItemRevenue, custerrs.MessageErr)
	GetTotalReportDateBetween(reportPayload DateBetween) ([]*entity.TotalItemRevenue, custerrs.MessageErr)
}
