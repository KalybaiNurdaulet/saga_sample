package saga

import (
	"context"
	"fmt"
	"log"
)

type Orchestrator struct {
	steps []Step
}

func NewOrchestrator(steps []Step) *Orchestrator {
	return &Orchestrator{steps: steps}
}

func (o *Orchestrator) Execute(ctx context.Context) error {
	var executedSteps []Step

	for _, step := range o.steps {
		log.Printf("[Saga] Executing step: %s", step.Name())
		if err := step.Execute(ctx); err != nil {
			log.Printf("[Saga] Step %s failed: %v. Starting compensation...", step.Name(), err)
			o.compensate(ctx, executedSteps)
			return fmt.Errorf("saga failed at step %s: %w", step.Name(), err)
		}
		executedSteps = append(executedSteps, step)
	}

	log.Println("[Saga] All steps completed successfully!")
	return nil
}

func (o *Orchestrator) compensate(ctx context.Context, steps []Step) {
	for i := len(steps) - 1; i >= 0; i-- {
		step := steps[i]
		log.Printf("[Saga] Compensating step: %s", step.Name())
		if err := step.Compensate(ctx); err != nil {
			
			log.Printf("[CRITICAL] Compensation failed for step %s: %v", step.Name(), err)
		}
	}
}