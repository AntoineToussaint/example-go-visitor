package business

import "context"

type Cache interface {
	GetStatistics(ctx context.Context) (*Statistics, error)
	SetStatistics(ctx context.Context, statistics *Statistics) error
}
