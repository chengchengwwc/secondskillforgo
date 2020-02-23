package services

import (
	"secondskillforgo/imooc_iris/datamodels"
	"secondskillforgo/imooc_iris/repositories"



)

type IOrderService interface {
	GetOrderById(int64)(*datamodels.Order,error)
	DeleteOrderById(int64) bool
	UpdateOrder(*datamodels.Order) error
	InsertOrder(*datamodels.Order) error
	GetAllOrder()([]*datamodels.Order,error)
	GetAllOrderInfo()(map[int]map[string]string,error)
}

type OrderService struct {
	OrderRepository repositories.IOrderRepository
}

func NewOrderService(OrderRepository repositories.IOrderRepository) IOrderService{
	return &OrderService{OrderRepository:OrderRepository}

}




