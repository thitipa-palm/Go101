package core

import "errors"

type OrderService interface { //primary port
	CreateOrder(order Order) error
}

type OrderServiceImpl struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService { //factory function
	return &OrderServiceImpl{repo: repo}
}

func (s *OrderServiceImpl) CreateOrder(order Order) error {
	//Business Logic function
	if order.Total <= 0 {
		return errors.New("Total have to be positive")
	}
	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}
