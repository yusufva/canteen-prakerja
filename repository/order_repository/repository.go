package order_repository

import (
	"canteen-prakerja/entity"
	"canteen-prakerja/pkg/custerrs"
)

type OrderRepository interface {
	GetAllOrder() ([]*entity.Item, custerrs.MessageErr)
	GetOrderById(orderId int) (*entity.Item, custerrs.MessageErr)
	CreateOrder(orderPayload []*entity.Item) (*entity.Item, custerrs.MessageErr)
	UpdateOrderById(orderId int, orderPayload *entity.Item) custerrs.MessageErr
	DeleteOrderById(orderId int) custerrs.MessageErr
}
