package main

import (
	"context"
	"fmt"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store: store}
}

func (s *service) CreateOrder(ctx context.Context) error {
	err := s.store.Create(ctx)
	if err != nil {
		//Handle the erro
		return err
	}

	fmt.Println("Order created successfully")
	return nil
}
