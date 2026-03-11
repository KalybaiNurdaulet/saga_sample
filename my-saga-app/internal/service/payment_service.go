package service

import (
	"context"
	"errors"
	"log"
)

type PaymentStep struct {
	Amount float64
}

func (p *PaymentStep) Name() string { return "Payment" }

func (p *PaymentStep) Execute(ctx context.Context) error {
	log.Printf("Processing payment of $%.2f...", p.Amount)
	if p.Amount > 1000 {
		return errors.New("insufficient funds or limit exceeded")
	}
	return nil
}

func (p *PaymentStep) Compensate(ctx context.Context) error {
	log.Println("Refunding payment...")
	return nil
}