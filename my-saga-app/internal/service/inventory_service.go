package service

import (
	"context"
	"errors"
	"log"
)

type InventoryStep struct {
	ItemID int
	Fail   bool
}

func (i *InventoryStep) Name() string { return "Inventory" }

func (i *InventoryStep) Execute(ctx context.Context) error {
	log.Printf("Reserving item %d in inventory...", i.ItemID)
	if i.Fail {
		return errors.New("item out of stock")
	}
	return nil
}

func (i *InventoryStep) Compensate(ctx context.Context) error {
	log.Printf("Releasing reservation for item %d...", i.ItemID)
	return nil
}