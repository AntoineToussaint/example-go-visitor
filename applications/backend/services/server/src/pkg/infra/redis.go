package infra

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/codefly-dev/go-grpc/base/pkg/business"
	"github.com/go-redis/redis/v8"
	"net/url"
)

type RedisCache struct {
	read  *redis.Client
	write *redis.Client
}

const key = "statistics"

func (c *RedisCache) GetStatistics(ctx context.Context) (*business.Statistics, error) {
	val, err := c.read.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	var model business.Statistics
	err = json.Unmarshal([]byte(val), &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (c *RedisCache) SetStatistics(ctx context.Context, statistics *business.Statistics) error {
	data, err := json.Marshal(statistics)
	if err != nil {
		return err
	}
	err = c.write.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewRedisClient(connectionString string) (*redis.Client, error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return nil, err
	}

	password := ""
	if u.User != nil {
		password, _ = u.User.Password()
	}

	return redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: password, // no password set
		DB:       0,        // use default DB
	}), nil
}

func NewRedisCache(connectionWrite string, connectionRead string) (*RedisCache, error) {
	read, err := NewRedisClient(connectionRead)
	if err != nil {
		return nil, err
	}
	write, err := NewRedisClient(connectionWrite)
	if err != nil {
		return nil, err
	}
	return &RedisCache{
		read:  read,
		write: write,
	}, nil
}
