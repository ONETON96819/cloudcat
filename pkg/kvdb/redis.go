package kvdb

import (
	"context"
	"time"

	goRedis "github.com/go-redis/redis/v8"
)

type redis struct {
	cli *goRedis.Client
}

func newRedis(cli *goRedis.Client) KvDb {
	return &redis{cli: cli}
}

func (r *redis) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.cli.Set(ctx, key, value, expiration).Err()
}

func (r *redis) Get(ctx context.Context, key string) (string, error) {
	ret, err := r.cli.Get(ctx, key).Result()
	if err != nil {
		if err == goRedis.Nil {
			return "", nil
		}
		return "", err
	}
	return ret, nil
}

func (r *redis) Has(ctx context.Context, key string) (bool, error) {
	ok, err := r.cli.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if ok == 1 {
		return true, nil
	}
	return false, nil
}

func (r *redis) Client() interface{} {
	return r.cli
}

func (r *redis) DbType() string {
	return "redis"
}
