package redis

import (
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/domain/models"
	custom "GenesisTask/pkg/errors"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
)

type rateTrait struct {
	Rate      float64   `json:"rate"`
	Timestamp time.Time `json:"timestamp"`
}

type redisCache struct {
	expirationTime time.Duration
	client         *redis.Client
}

func NewRedisCache(host string, db int, exp time.Duration) application.Cache {
	return &redisCache{client: getRedisClient(host, db), expirationTime: exp}
}

func getRedisClient(host string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: host,
		DB:   db,
	})
}

func (r *redisCache) AddRateToCache(rate models.CurrencyRate) {
	pair := rate.GetCurrencyPair()
	rateAssets := pair.ToString()
	rateJson, _ := json.Marshal(rateTrait{rate.GetPrice(), rate.GetTimestamp()})
	r.client.Set(context.Background(), rateAssets, rateJson, r.expirationTime)
}

func (r *redisCache) GetRateFromCache(pair models.CurrencyPair) (models.CurrencyRate, error) {
	rateAssets := pair.ToString()
	rateBytes, err := r.client.Get(context.Background(), rateAssets).Bytes()
	if errors.Is(err, redis.Nil) {
		return *models.NewCurrencyRate(pair, -1, time.Now()), custom.ErrNotPresentInCache
	}
	if err != nil {
		return *models.NewCurrencyRate(pair, -1, time.Now()), err
	}

	var result rateTrait
	err = json.Unmarshal(rateBytes, &result)
	if err != nil {
		return *models.NewCurrencyRate(pair, -1, time.Now()), err
	}
	return *models.NewCurrencyRate(pair, result.Rate, result.Timestamp), nil
}
