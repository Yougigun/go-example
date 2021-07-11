package redis

import (
	"context"
	"testing"
	"themis/component/cache"
	"time"
	// cpredis "themis/component/cache/provider/redis"
)

func TestProvider(t *testing.T) {
	ctx := context.TODO()

	conf := cache.Config{
		Namespace:   "test",
		GlobalExpireTime: time.Second * 5,
	}

	providerRedisConfig := Config{
		Addrs:        "redis.fe-cluster.devfdg.net:6379",
		Username:     "",
		Password:     "",
		DB:           0,
		PoolSize:     30,
		MinIdleConns: 10,
		IdleTimeout:  time.Second * 30,
	}
	CacheManager := cache.NewManager(ctx, &conf, ProviderName, &providerRedisConfig)
	
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	
	//Set cache
	mykey := "mykey"
	mycache := User{Name: "gary", Age: 18}
	expire := time.Second * 10 // 0: use default GlobalExpireTime
	err := CacheManager.Set(ctx, mykey, mycache, expire) 
	if err != nil {
		panic(err)
	}

	//Get cache
	wanna := User{}
	if ok,err:=CacheManager.Get(ctx, mykey, &wanna);err!=nil{
		panic(err)
	}else if !ok&&err==nil{
		t.Log("cache miss")
	}else{
		t.Log("cache hit")
		t.Log(wanna)
		// output:{gary 18}
	}

	//Delete cache
	if err :=CacheManager.Delete(ctx,mykey);err!=nil{
		panic(err)
	}

	//Close
	CacheManager.Close()
}
