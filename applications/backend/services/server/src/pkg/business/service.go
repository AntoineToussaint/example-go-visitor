package business

import (
	"context"
	"fmt"
	"github.com/codefly-dev/core/wool"
	"github.com/codefly-dev/go-grpc/base/pkg/gen"
)

type Service struct {
	store Store
	cache Cache
}

func (c *Service) Visit(ctx context.Context) (int, error) {
	if c.store == nil {
		return 0, fmt.Errorf("store is not set")
	}
	visit, err := c.store.Visit(ctx)
	if err != nil {
		return 0, err
	}
	return visit, nil
}

type Statistics struct {
	Total  int
	Visits []*gen.DailyVisitStatistics
}

func (c *Service) GetStatistics(ctx context.Context) (*Statistics, error) {
	w := wool.Get(ctx).In("GetStatistics")
	if c.cache != nil {
		stats, err := c.cache.GetStatistics(ctx)
		if err != nil {
			return nil, err
		}
		if stats != nil {
			w.Info("using cached value")
			return stats, nil
		}
	}
	stats, err := c.store.GetStatistics(ctx)
	if err != nil {
		return nil, err
	}
	if c.cache != nil {
		if err := c.cache.SetStatistics(ctx, stats); err != nil {
			return nil, err
		}
	}
	return stats, nil
}

func (c *Service) WithStore(store Store) {
	c.store = store
}

func (c *Service) WithCache(cache Cache) {
	c.cache = cache
}
