package redis

import (
	// Go Internal Packages
	"context"
	"errors"
	"time"

	// External Packages
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisHandler struct {
	masterClient *redis.Client
	slaveClient  *redis.Client
	log          *zap.Logger
}

type RedisElement struct {
	Key    string        // Default : ""
	Val    any           // Default : nil
	Expiry time.Duration // Default : 0
}

func (e RedisElement) value() any {
	if e.Val != nil {
		return e.Val
	}
	return "1"
}

// NewRedisHandler creates a new RedisHandler instance consisting of master and slave clients
func NewRedisHandler(masterClient *redis.Client, slaveClient *redis.Client, log *zap.Logger) *RedisHandler {
	return &RedisHandler{masterClient, slaveClient, log}
}

// Close closes the master and slave clients
func (r *RedisHandler) Close() {
	r.masterClient.Close()
	r.slaveClient.Close()
}

// Ping pings the master and slave clients to check the connection
func (r *RedisHandler) Ping(ctx context.Context) error {
	_, sPingErr := r.slaveClient.Ping(ctx).Result()

	if sPingErr != nil {
		return sPingErr
	}

	_, mPingErr := r.masterClient.Ping(ctx).Result()

	if mPingErr != nil {
		return mPingErr
	}

	return nil
}

// Exists checks if a key exists in the redis
func (r *RedisHandler) Exists(ctx context.Context, key string) (bool, error) {
	res, err := r.slaveClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if res != 1 {
		return false, nil
	}
	return true, nil
}

// Insert inserts a key-value pair to redis
func (r *RedisHandler) Insert(ctx context.Context, element RedisElement) error {
	if element.Key == "" {
		return errors.New("invalid operation : empty key for redis set operation")
	}
	return r.masterClient.Set(ctx, element.Key, element.value(), element.Expiry).Err()
}

// Get returns the value of the key from the redis
func (r *RedisHandler) Get(ctx context.Context, key string) (string, error) {
	return r.slaveClient.Get(ctx, key).Result()
}

// DeleteKey deletes a key from redis
func (r *RedisHandler) DeleteKey(ctx context.Context, key string) (int64, error) {
	return r.masterClient.Del(ctx, key).Result()
}
