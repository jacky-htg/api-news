package libraries

/*
import (
	"time"

	"github.com/jacky-htg/api-news/config"
	r "gopkg.in/redis.v5"
)

const PREFIX_REDIS = "news_api:"

var client = r.NewClient(&r.Options{
	Addr:     config.GetString("database.redis.address"),
	Password: "",
})

func RedisSet(key string, val interface{}, duration time.Duration) error {
	key = PREFIX_REDIS + key
	if err := client.SetNX(key, val, duration).Err(); err != nil {
		return err
	}

	return nil
}

func RedisGet(key string) ([]byte, error) {
	key = PREFIX_REDIS + key
	getting, err := client.Get(key).Bytes()
	if err != nil {
		return []byte{}, err
	}

	return getting, nil
}

func RedisExists(key string) bool {
	key = PREFIX_REDIS + key
	if exists := client.Exists(key).Val(); exists {
		return true
	}

	return false
}

func RedisDelete(key string) (bool, error) {
	key = PREFIX_REDIS + key
	if err := client.Del(key).Err(); err != nil {
		return false, err
	}

	return true, nil
}

func RedisHashGet(key string, field string) (string, error) {
	key = PREFIX_REDIS + key
	cache := client.HGet(key, field)
	if cache.Err() != nil {
		return "", cache.Err()
	}

	return cache.Val(), nil
}

func RedisHashSet(key string, field string, val interface{}) error {
	key = PREFIX_REDIS + key
	if err := client.HSet(key, field, val).Err(); err != nil {
		return err
	}

	return nil
}

func RedisHashExists(key string, field string) bool {
	key = PREFIX_REDIS + key
	if exists := client.HExists(key, field).Val(); exists {
		return true
	}

	return false
}
*/
