package saga

import "context"

type Step interface {
	Name() string
	Execute(ctx context.Context) error
	Compensate(ctx context.Context) error
}