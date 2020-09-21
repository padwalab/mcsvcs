package custom

import "context"

type Service interface {
	TestStatus(ctx context.Context) (int, error)
}
