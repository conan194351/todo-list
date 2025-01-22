package redis

import (
	"context"
	"fmt"
	"github.com/conan194351/todo-list.git/pkg/logger"
	"time"

	"github.com/conan194351/todo-list.git/internal/config"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
	logger logger.Logger
}

var Clt *Client

func NewClient() *Client {
	cnf := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cnf.Redis.Host, cnf.Redis.Port),
		Password: cnf.Redis.Password,
		DB:       0,
	})
	Clt = &Client{
		client: client,
		logger: logger.NewZapLogger("redis", true),
	}
	return Clt
}

func (r *Client) Save(key string, value interface{}, duration time.Duration) error {
	ctx := context.Background()
	err := r.client.Set(ctx, key, value, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Client) Delete(key string) error {
	ctx := context.Background()
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Client) Get(key string) (string, error) {
	ctx := context.Background()
	tokensJSON, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return tokensJSON, nil
}

func (r *Client) Publish(ctx context.Context, channel string, message interface{}) error {
	err := r.client.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Client) Subscribe(ctx context.Context, channel ...string) *redis.PubSub {
	return r.client.Subscribe(ctx, channel...)
}
