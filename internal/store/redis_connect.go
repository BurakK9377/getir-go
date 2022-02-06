package store

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"net/url"
	"os"
	"strings"
	"time"
)

type RedisConfiguration struct {
	RedisUrl string
}

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

// Initialize GetClient get the redis client
func Initialize() *redisClient {
	config := GetRedisConfiguration()
	var resolvedURL = config.RedisUrl
	var password = ""

	if !strings.Contains(resolvedURL, "localhost") {
		parsedURL, _ := url.Parse(resolvedURL)
		password, _ = parsedURL.User.Password()
		resolvedURL = parsedURL.Host
	}

	c := redis.NewClient(&redis.Options{
		Addr:     resolvedURL,
		Password: password,
		DB:       0, // use default DB
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	client.c = c
	return client
}

// GetRedisConfiguration method basically populate configuration information from .env and return Configuration model
func GetRedisConfiguration() RedisConfiguration {
	_ = godotenv.Load(".env")

	configuration := RedisConfiguration{
		os.Getenv("REDIS_URL"),
	}

	return configuration
}

//GetKey get key
func (client *redisClient) GetKey(key string, src interface{}) error {
	val, err := client.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

//SetKey set key
func (client *redisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
