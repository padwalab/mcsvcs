package edil

import (
	"context"

	"github.com/padwalab/mcsvcs/internal"
)

// Service interface
type Service interface {
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	ServiceStatus(ctx context.Context) (int, error)
}
