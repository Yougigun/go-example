package redis

import (
	"context"
	"fmt"
	"strings"
	"go-example/cache"
	"time"

	"github.com/go-redis/redis/v8"
	// "themis/component/redis"
)

func init() {
	cache.Register(ProviderName, newProvider())
}

const ProviderName cache.ProviderName = "redis"

type provider struct {
	WriteClient *redis.ClusterClient
	ReadClient  *redis.ClusterClient
}

// type Handler struct {
// 	WriteClient *redis.Client
// 	ReadClient  *redis.Client
// }
type Config struct {
	Addrs        string        `yaml:"addrs" json:"addrs"`
	Username     string        `yaml:"username" json:"username"`
	Password     string        `yaml:"password" json:"password"`
	PoolSize     int           `yaml:"pool_size" json:"pool_size"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" json:"idle_timeout"`
	MinIdleConns int           `yaml:"min_idle_conns" json:"min_idle_conns"`
	DB           int           `yaml:"db" json:"db"`
}

func newProvider() *provider {
	return &provider{}
}
func (p *provider) Set(
	ctx context.Context, key, value string, expire time.Duration) error {
	err := p.WriteClient.Set(ctx, key, value, expire).Err()
	return err
}

// if err == nil && bool == false,then cache miss.
func (p *provider) Get(ctx context.Context, key string) (string, bool, error) {
	value, err := p.ReadClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return value, false, nil
	} else if err != nil {
		return value, false, err
	}
	return value, true, nil
}

func (p *provider) Delete(ctx context.Context, keys ...string) error {
	if _, err := p.WriteClient.Del(ctx, keys...).Result(); err != nil {
		return err
	} else {
		// success delete
		return nil
	}
}
func (p *provider) Close() {
	p.WriteClient.Close()
	p.ReadClient.Close()
}

func (p *provider) Init(ctx context.Context, config interface{}) error {
	c := config.(*Config)
	readClient := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs:              strings.Split(c.Addrs, ","),
			ReadOnly:           true,
			RouteRandomly:      true,
			Username:           c.Username,
			Password:           c.Password,
			PoolSize:           c.PoolSize,
			IdleTimeout:        c.IdleTimeout,
			MinIdleConns:       c.MinIdleConns,
			IdleCheckFrequency: c.IdleTimeout,
		},
	)
	writeClient := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs:              strings.Split(c.Addrs, ","),
			Username:           c.Username,
			Password:           c.Password,
			PoolSize:           c.PoolSize,
			IdleTimeout:        c.IdleTimeout,
			MinIdleConns:       c.MinIdleConns,
			IdleCheckFrequency: c.IdleTimeout,
		},
	)
	val, err := writeClient.Ping(ctx).Result()
	if val != "PONG" {
		return fmt.Errorf("failed to create write Redis client: %s, err: %w", val, err)
	}
	val, err = readClient.Ping(ctx).Result()
	if val != "PONG" {
		return fmt.Errorf("failed to create write Redis client: %s, err: %w", val, err)
	}
	*p = provider{
		WriteClient: writeClient,
		ReadClient:  readClient}
	return nil
}
