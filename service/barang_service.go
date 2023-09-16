package service

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/pkg/helpers"
	"canteen-prakerja/repository/barang_repository"
	"net/http"
	"time"
)

type BarangService interface {
	GetAllBarang() (*dto.GetBarangResponse, custerrs.MessageErr)
	GetBarangById(barangId int) (*dto.BarangResponse, custerrs.MessageErr)
	CreateBarang(barangPayload dto.NewBarangRequest) (*dto.NewBarangResponse, custerrs.MessageErr)
	UpdateBarangById(barangId int, barangPayload dto.NewBarangRequest) (*dto.NewBarangResponse, custerrs.MessageErr)
	DeleteBarangById(barangId int) (*dto.NewBarangResponse, custerrs.MessageErr)
}

func NewBarangService(barangRepo barang_repository.BarangRepository) BarangService {
	return &barangService{
		barangRepo: barangRepo,
	}
}

type barangService struct {
	barangRepo barang_repository.BarangRepository
}

func (b *barangService) GetAllBarang() (*dto.GetBarangResponse, custerrs.MessageErr) {
	barangs, err := b.barangRepo.GetAllBarang()

	if err != nil {
		return nil, err
	}

	barangResponse := []dto.BarangResponse{}

	for _, eachBarang := range barangs {
		barangResponse = append(barangResponse, eachBarang.EntityToBarangResponseDto())
	}

	response := dto.GetBarangResponse{
		Result:     "success",
		Statuscode: http.StatusOK,
		Data:       barangResponse,
	}

	return &response, nil
}

func (b *barangService) GetBarangById(barangId int) (*dto.BarangResponse, custerrs.MessageErr) {
	result, err := b.barangRepo.GetBarangById(barangId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToBarangResponseDto()

	return &response, nil
}

func (b *barangService) CreateBarang(barangPayload dto.NewBarangRequest) (*dto.NewBarangResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(barangPayload)

	if err != nil {
		return nil, err
	}

	barangRequest := entity.Barang{
		Barang:    barangPayload.Barang,
		HargaBeli: barangPayload.HargaBeli,
		HargaJual: barangPayload.HargaJual,
	}

	_, err = b.barangRepo.CreateBarang(&barangRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewBarangResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "barang has been successfully created",
	}

	return &response, nil
}

func (b *barangService) UpdateBarangById(barangId int, barangPayload dto.NewBarangRequest) (*dto.NewBarangResponse, custerrs.MessageErr) {
	err := helpers.ValidateStruct(barangPayload)

	if err != nil {
		return nil, err
	}

	payload := entity.Barang{
		ID:        barangId,
		Barang:    barangPayload.Barang,
		HargaBeli: barangPayload.HargaBeli,
		HargaJual: barangPayload.HargaJual,
		UpdatedAt: time.Now(),
	}

	err = b.barangRepo.UpdateBarangById(&payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewBarangResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "barang has been updated",
	}

	return &response, nil
}

func (b *barangService) DeleteBarangById(barangId int) (*dto.NewBarangResponse, custerrs.MessageErr) {
	err := b.barangRepo.DeleteBarangById(barangId)

	if err != nil {
		return nil, err
	}

	response := dto.NewBarangResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "barang has been successfully deleted",
	}

	return &response, nil
}
