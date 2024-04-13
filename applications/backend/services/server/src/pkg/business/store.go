package business

import "context"

type Store interface {
	Visit(ctx context.Context) (int, error)
	GetStatistics(ctx context.Context) (*Statistics, error)
}
