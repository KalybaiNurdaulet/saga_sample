package service

import (
	"context"
	"log"
)

type ShippingStep struct {
	Address string
}

func (s *ShippingStep) Name() string { return "Shipping" }

func (s *ShippingStep) Execute(ctx context.Context) error {
	log.Printf("Creating shipping label for: %s", s.Address)
	return nil
}

func (s *ShippingStep) Compensate(ctx context.Context) error {
	log.Println("Cancelling shipping label...")
	return nil
}