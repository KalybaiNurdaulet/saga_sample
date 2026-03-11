package service

import (
	"context"
	"fmt"
	"my-saga-app/internal/saga"
)


type CheckoutRequest struct {
	OrderID    string
	CustomerID string
	ItemID     int
	Amount     float64
	Address    string
}


type CheckoutSagaService struct {
}

func NewCheckoutSagaService() *CheckoutSagaService {
	return &CheckoutSagaService{}
}


func (s *CheckoutSagaService) Run(ctx context.Context, req CheckoutRequest) error {

	steps := []saga.Step{
		&PaymentStep{
			Amount: req.Amount,
		},
		&InventoryStep{
			ItemID: req.ItemID,
			
		},
		&ShippingStep{
			Address: req.Address,
		},
	}


	orchestrator := saga.NewOrchestrator(steps)


	if err := orchestrator.Execute(ctx); err != nil {
		
		return fmt.Errorf("checkout saga failed for order %s: %w", req.OrderID, err)
	}

	return nil
}