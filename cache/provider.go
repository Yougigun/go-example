package cache

import (
	"context"
	"time"
)

var providers = map[ProviderName]provider{}

type ProviderName string

type provider interface {
	Set(context.Context, string, string, time.Duration) error
	Get(context.Context, string) (string, bool, error)
	Delete(ct context.Context, keys ...string) error
	Init(ctx context.Context, config interface{}) error
	Close()
}

func Register(providerName ProviderName, provider provider) {
	providers[providerName] = provider
}
