package dto

type ReportResponse struct {
	Barang         string `json:"barang"`
	TotalTerjual   int    `json:"total_terjual"`
	TotalModal     int    `json:"total_modal"`
	TotalHargaJual int    `json:"total_harga_jual"`
	Laba           int    `json:"laba"`
}

type GetReportResponse struct {
	Result     string             `json:"result"`
	StatusCode int                `json:"statuscode"`
	Message    string             `json:"message"`
	Data       []ReportResponse   `json:"data"`
	Total      []TotalSumResponse `json:"sum_data"`
}

type TotalSumResponse struct {
	SumTerjual   int `json:"sum_terjual"`
	SumModal     int `json:"sum_modal"`
	SumHargaJual int `json:"sum_hargajual"`
	SumLaba      int `json:"sum_laba"`
}

type SingleDateReportRequest struct {
	Date string `json:"date"`
}

type DateRangeReportRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
