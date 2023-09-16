package service

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/pkg/helpers"
	"canteen-prakerja/repository/transaksi_repository"
	"net/http"
)

type TransaksiService interface {
	GetAllTransaksi() (*dto.GetTransaksiResponse, custerrs.MessageErr)
	GetTransaksiById(transaksiId int) (*dto.TransaksiResponse, custerrs.MessageErr)
	GetTransaksiDateBetween(transaksiPayload *dto.TransaksiDateBetweenRequest) (*dto.GetTransaksiResponse, custerrs.MessageErr)
	CreateTransaksi(transaksiPayload *dto.NewTransaksiRequest) (*dto.GetTransaksiResponse, custerrs.MessageErr)
	// UpdateTransaksiById(transaksiId int, transaksiPayload *entity.Transaksi) custerrs.MessageErr
	DeleteTransaksiById(transaksiId int) (*dto.NewTransaksiResponse, custerrs.MessageErr)
}

type transaksiService struct {
	transaksiRepo transaksi_repository.TransaksiRepository
}

func NewTransaksiService(transaksiRepo transaksi_repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepo: transaksiRepo,
	}
}

func (t *transaksiService) GetAllTransaksi() (*dto.GetTransaksiResponse, custerrs.MessageErr) {
	txs, err := t.transaksiRepo.GetAllTransaksi()

	if err != nil {
		return nil, err
	}

	txResponse := []dto.TransaksiResponse{}

	for _, eachTx := range txs {
		items := []dto.ItemResponse{}
		for _, eachItems := range eachTx.Items {
			items = append(items, eachItems.EntityToItemResponseDto())
		}
		txResponse = append(txResponse, dto.TransaksiResponse{
			ID:         eachTx.ID,
			Tanggal:    eachTx.Tanggal,
			Items:      items,
			TotalHarga: eachTx.TotalHarga,
			CreatedAt:  eachTx.CreatedAt,
			UpdatedAt:  eachTx.UpdatedAt,
		})
	}

	response := dto.GetTransaksiResponse{
		Result:     "succes",
		StatusCode: http.StatusOK,
		Message:    "transaction data has been sent successfully",
		Data:       txResponse,
	}

	return &response, nil
}

func (t *transaksiService) GetTransaksiById(transaksiId int) (*dto.TransaksiResponse, custerrs.MessageErr) {
	res, err := t.transaksiRepo.GetTransaksiById(transaksiId)

	if err != nil {
		return nil, err
	}

	items := []dto.ItemResponse{}
	for _, eachItems := range res.Items {
		items = append(items, eachItems.EntityToItemResponseDto())
	}

	response := dto.TransaksiResponse{
		ID:         res.ID,
		Tanggal:    res.Tanggal,
		Items:      items,
		TotalHarga: res.TotalHarga,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
	}

	return &response, nil
}

func (t *transaksiService) GetTransaksiDateBetween(transaksiPayload *dto.TransaksiDateBetweenRequest) (*dto.GetTransaksiResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(transaksiPayload)

	if err != nil {
		return nil, err
	}

	payload := transaksi_repository.DateBetween{
		StartDate: transaksiPayload.StartDate,
		EndDate:   transaksiPayload.EndDate,
	}
	txs, err := t.transaksiRepo.GetTransaksiDateBetween(payload)

	if err != nil {
		return nil, err
	}

	txResponse := []dto.TransaksiResponse{}

	for _, eachTx := range txs {
		items := []dto.ItemResponse{}
		for _, eachItems := range eachTx.Items {
			items = append(items, eachItems.EntityToItemResponseDto())
		}
		txResponse = append(txResponse, dto.TransaksiResponse{
			ID:         eachTx.ID,
			Tanggal:    eachTx.Tanggal,
			Items:      items,
			TotalHarga: eachTx.TotalHarga,
			CreatedAt:  eachTx.CreatedAt,
			UpdatedAt:  eachTx.UpdatedAt,
		})
	}

	response := dto.GetTransaksiResponse{
		Result:     "succes",
		StatusCode: http.StatusOK,
		Message:    "transaction data has been sent successfully",
		Data:       txResponse,
	}

	return &response, nil
}

func (t *transaksiService) CreateTransaksi(transaksiPayload *dto.NewTransaksiRequest) (*dto.GetTransaksiResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(transaksiPayload)

	if err != nil {
		return nil, err
	}

	items := []entity.Item{} //make([]entity.Item, len(transaksiPayload.Items))
	for _, eachItem := range transaksiPayload.Items {
		items = append(items, entity.Item{
			Barang:     eachItem.Barang,
			Jumlah:     eachItem.Jumlah,
			HargaJual:  eachItem.HargaJual,
			HargaBeli:  eachItem.HargaBeli,
			Laba:       (eachItem.HargaJual - eachItem.HargaBeli) * eachItem.Jumlah,
			TotalHarga: eachItem.TotalHarga,
		})
	}

	txRequest := &entity.Transaksi{
		Tanggal:    transaksiPayload.Tanggal,
		Items:      items,
		TotalHarga: transaksiPayload.TotalHarga,
	}

	result, err := t.transaksiRepo.CreateTransaksi(txRequest)

	if err != nil {
		return nil, err
	}

	dtoResult := []dto.TransaksiResponse{}
	dtoResult = append(dtoResult, result.EntityToTransaksiResponseDTO())

	response := dto.GetTransaksiResponse{
		Result:     "success",
		Message:    "transaction has been successfully created",
		StatusCode: http.StatusCreated,
		Data:       dtoResult,
	}

	return &response, nil
}

// UpdateTransaksiById(transaksiId int, transaksiPayload *entity.Transaksi) custerrs.MessageErr

func (t *transaksiService) DeleteTransaksiById(transaksiId int) (*dto.NewTransaksiResponse, custerrs.MessageErr) {
	err := t.transaksiRepo.DeleteTransaksiById(transaksiId)

	if err != nil {
		return nil, err
	}

	response := dto.NewTransaksiResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "transaction has been successfully deleted",
	}

	return &response, nil
}
