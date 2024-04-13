package main

import (
	"context"
	"github.com/codefly-dev/core/wool"
	"github.com/codefly-dev/go-grpc/base/pkg/adapters"
	"github.com/codefly-dev/go-grpc/base/pkg/business"
	"github.com/codefly-dev/go-grpc/base/pkg/infra"
	codefly "github.com/codefly-dev/sdk-go"
)

func init() {
	WithWork(doWork)
}

func doWork(ctx context.Context) (Clean, error) {
	w := wool.Get(ctx)
	service := &business.Service{}

	// Postgres
	conn, err := codefly.For(ctx).Service("store").Secret("postgres", "connection")
	if err != nil {
		return nil, err
	}
	store, err := infra.NewPostgresStore(conn)
	if err != nil {
		return nil, err
	}
	service.WithStore(store)

	connectionRead, err := codefly.For(ctx).Service("cache").Secret("read", "connection")
	if err != nil {
		return nil, err
	}
	connectionWrite, err := codefly.For(ctx).Service("cache").Secret("write", "connection")
	if err != nil {
		return nil, err
	}
	w.Info("configuration",
		wool.Field("postgres connection", conn),
		wool.Field("redis read connection", connectionRead),
		wool.Field("redis write connection", connectionWrite))
	// Redis
	cache, err := infra.NewRedisCache(connectionWrite, connectionRead)
	if err != nil {
		return nil, err
	}
	service.WithCache(cache)

	adapters.SetService(service)
	return func() {}, nil
}
