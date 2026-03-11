package main

import (
	"context"
	"log"
	"time"

	"my-saga-app/internal/saga"
	"my-saga-app/internal/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("--- Starting Successful Checkout Saga ---")
	runCheckout(ctx, 100.0, 42, "Wall St 1", false)

	log.Println("\n--- Starting Failing Checkout Saga (Inventory Fail) ---")
	runCheckout(ctx, 50.0, 99, "Broadway 10", true)
}

func runCheckout(ctx context.Context, price float64, itemID int, addr string, shouldFailInventory bool) {
	
	steps := []saga.Step{
		&service.PaymentStep{Amount: price},
		&service.InventoryStep{ItemID: itemID, Fail: shouldFailInventory},
		&service.ShippingStep{Address: addr},
	}


	orchestrator := saga.NewOrchestrator(steps)

	// Запуск
	if err := orchestrator.Execute(ctx); err != nil {
		log.Printf("Checkout flow finished with error: %v", err)
	} else {
		log.Println("Checkout flow finished successfully!")
	}
}