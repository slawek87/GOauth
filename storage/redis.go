package storage

import (
	"github.com/go-redis/redis"
)

func RedisDB() (*redis.Client, error){
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	_, err := db.Ping().Result()

	if err != nil {
		return nil, err
	}

	return db, nil
}