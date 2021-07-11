package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Manager struct {
	namespace   string
	globalExpireTime time.Duration
	provider    provider
}

type Config struct {
	Namespace   string        `yaml:"namespace" json:"namespace"`
	GlobalExpireTime time.Duration `yaml:"global_expire_time" json:"global_expire_time"`
}

// Get the provider from "provider package"
func NewManager(ctx context.Context, conf *Config, provider ProviderName, providerConfig interface{}) *Manager {
	manager := &Manager{}
	if p, ok := providers[provider]; !ok {
		panic(fmt.Sprintf("%v does not register.", provider))
	} else {
		manager = &Manager{
			namespace:   conf.Namespace,
			provider:    p,
			globalExpireTime: conf.GlobalExpireTime,
		}
	}
	if err := manager.provider.Init(ctx, providerConfig); err != nil {
		panic(err)
	}
	return manager
}

// Set the cache if expire is 0,then use globalExpireTime
func (m *Manager) Set(ctx context.Context, key string, value interface{}, expire time.Duration) (err error) {
	key = m.generateKey(key)
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if expire == 0 {
		expire = m.globalExpireTime
	}
	if err:=m.provider.Set(ctx, key, string(b), expire);err!=nil{
		return err
	}
	return nil
}

// if bool is true, which mean cache hit.
func (m *Manager) Get(ctx context.Context, key string, obj interface{}) (bool, error) {
	key = m.generateKey(key)
	value, ok, err := m.provider.Get(ctx, key)
	if err != nil||!ok{
		return ok, err
	}
	if err := json.Unmarshal([]byte(value), obj); err != nil {
		return false, err
	}
	return ok, err
}
func (m *Manager) Delete(ctx context.Context, keys ...string) error {
	tempKeys := make([]string, len(keys))
	for i, key := range keys {
		tempKeys[i] = m.generateKey(key)
	}
	if err := m.provider.Delete(ctx, tempKeys...); err != nil {
		return err
	}
	return nil
}

func (m *Manager) Close(ctx context.Context) {
	m.provider.Close()
}

func (m *Manager) generateKey(key string) string {
	return m.namespace + "-" + key
}
