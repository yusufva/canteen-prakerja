package service

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/pkg/helpers"
	"canteen-prakerja/repository/report_repository"
	"net/http"
)

type ReportService interface {
	GetReportDateBetween(reportPayload *dto.DateRangeReportRequest) (*dto.GetReportResponse, custerrs.MessageErr)
	GetTotalReportDateBetween(reportPayload *dto.DateRangeReportRequest) (*dto.GetReportResponse, custerrs.MessageErr)
}

type reportService struct {
	reportRepo report_repository.ReportRepository
}

func NewReportService(reportRepo report_repository.ReportRepository) ReportService {
	return &reportService{
		reportRepo: reportRepo,
	}
}

func (rs *reportService) GetReportDateBetween(reportPayload *dto.DateRangeReportRequest) (*dto.GetReportResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(reportPayload)

	if err != nil {
		return nil, err
	}

	payload := report_repository.DateBetween{
		StartDate: reportPayload.StartDate,
		EndDate:   reportPayload.EndDate,
	}

	report, err := rs.reportRepo.GetReportDateBetween(payload)

	if err != nil {
		return nil, err
	}

	total, err := rs.reportRepo.GetTotalReportDateBetween(payload)

	if err != nil {
		return nil, err
	}

	reportResponse := []dto.ReportResponse{}

	for _, eachReport := range report {
		reportResponse = append(reportResponse, eachReport.EntityToReportResponseDTO())
	}

	totalResponse := []dto.TotalSumResponse{}

	for _, eachTotal := range total {
		totalResponse = append(totalResponse, eachTotal.EntityToTotalSumResponseDTO())
	}

	// totalResponse = append(totalResponse, dto.TotalSumResponse{
	// 	SumTerjual:   total.JumlahTot,
	// 	SumModal:     total.HargaBeliTot,
	// 	SumHargaJual: total.HargaJualTot,
	// 	SumLaba:      total.Laba,
	// })

	// totalResponse = append(totalResponse, dto.TotalSumResponse{
	// 	SumTerjual:   total[0].JumlahTot,
	// 	SumModal:     total[0].HargaBeliTot,
	// 	SumHargaJual: total[0].HargaJualTot,
	// 	SumLaba:      total[0].Laba,
	// })

	response := dto.GetReportResponse{
		Result:     "success",
		Message:    "report has been successfully send",
		StatusCode: http.StatusOK,
		Data:       reportResponse,
		Total:      totalResponse,
	}

	return &response, nil
}

func (rs *reportService) GetTotalReportDateBetween(reportPayload *dto.DateRangeReportRequest) (*dto.GetReportResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(reportPayload)

	if err != nil {
		return nil, err
	}

	payload := report_repository.DateBetween{
		StartDate: reportPayload.StartDate,
		EndDate:   reportPayload.EndDate,
	}

	total, err := rs.reportRepo.GetTotalReportDateBetween(payload)

	if err != nil {
		return nil, err
	}

	totalResponse := []dto.TotalSumResponse{}
	totalResponse = append(totalResponse, total[0].EntityToTotalSumResponseDTO())

	response := dto.GetReportResponse{
		Result:     "success",
		Message:    "report has been successfully send",
		StatusCode: http.StatusOK,
		// Data:       reportResponse,
		Total: totalResponse,
	}

	return &response, nil
}
