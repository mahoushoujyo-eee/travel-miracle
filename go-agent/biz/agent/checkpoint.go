package agent

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CheckPointStore interface {
    Set(ctx context.Context, key string, value []byte) error
    Get(ctx context.Context, key string) ([]byte, bool, error)
}

type InMemoryCheckPointStore struct {
    data map[string][]byte
}

func NewInMemoryStore() *InMemoryCheckPointStore {
    return &InMemoryCheckPointStore{
        data: make(map[string][]byte),
    }
}

func (s *InMemoryCheckPointStore) Set(ctx context.Context, key string, value []byte) error {
    s.data[key] = value
    return nil
}

func (s *InMemoryCheckPointStore) Get(ctx context.Context, key string) ([]byte, bool, error) {
    value, ok := s.data[key]
    return value, ok, nil
}

type RedisCheckPointStore struct {
    client *redis.Client
}

func NewRedisCheckPointStore(client *redis.Client) *RedisCheckPointStore {
    return &RedisCheckPointStore{
        client: client,
    }
}

func (s *RedisCheckPointStore) Set(ctx context.Context, key string, value []byte) error {
    return s.client.Set(ctx, key, value, 0).Err()
}

func (s *RedisCheckPointStore) Get(ctx context.Context, key string) ([]byte, bool, error) {
    val, err := s.client.Get(ctx, key).Bytes()
    if err != nil {
        if err == redis.Nil {
            return nil, false, nil
        }
        return nil, false, err
    }
    return val, true, nil
}