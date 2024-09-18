package main

import (
	"context"
)

type store struct {
	// add mongodb instance
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	// if something goes wrong return an error
	return nil
}
