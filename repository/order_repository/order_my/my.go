package order_my

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"

	"gorm.io/gorm"
)

type orderMY struct {
	db *gorm.DB
}

func (o *orderMY) GetAllOrder() ([]*entity.Item, custerrs.MessageErr) {
	var order []*entity.Item
	err := o.db.Find(&order).Error

	if err != nil {
		return nil, custerrs.NewInternalServerError("error while getting order data")
	}

	return order, nil
}

func (o *orderMY) GetOrderById(orderId int) (*entity.Item, custerrs.MessageErr) {
	var order entity.Item
	result := o.db.First(&order, orderId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, custerrs.NewNotFoundError("order not found")
		}
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	return &order, nil
}

func (o *orderMY) CreateOrder(orderPayload []*entity.Item) ([]*entity.Item, custerrs.MessageErr) {
	result := o.db.Create(orderPayload)

	if result.Error != nil {
		return nil, custerrs.NewInternalServerError("something went wrong")
	}

	row := result.Row()
	var orders []*entity.Item
	row.Scan(row, &orders)

	return orders, nil
}

func (o *orderMY) UpdateOrderById(orderId int, orderPayload *entity.Item) custerrs.MessageErr {
	err := o.db.Model(orderPayload).Updates(entity.Item{TransaksiID: orderPayload.TransaksiID, Barang: orderPayload.Barang, HargaJual: orderPayload.HargaJual, HargaBeli: orderPayload.HargaBeli}).Error

	if err != nil {
		return custerrs.NewInternalServerError("error while updating orders")
	}

	return nil
}

func (o *orderMY) DeleteOrderById(orderId int) custerrs.MessageErr {
	var order entity.Item
	result := o.db.First(&order, orderId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return custerrs.NewNotFoundError("order not found")
		}
		return custerrs.NewInternalServerError("something went wrong")
	}

	err := o.db.Delete(&order).Error

	if err != nil {
		return custerrs.NewInternalServerError("something went wrong")
	}

	return nil
}
