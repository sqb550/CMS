package utils

import (
	"CMS/app/models"
	"CMS/config/database"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:8080",
		Password: "",
		DB:       0,
	})
}

func GetPostCacheKey(PostID uint) string {
	return fmt.Sprintf("post:%d", PostID)
}

// 从缓存获取用户
func GetPostFromCache(PostID uint) (*models.Post, bool, error) {
	key := GetPostCacheKey(PostID)
	values, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, false, nil
		}
		return nil, false, err
	}
	var post models.Post
	err = json.Unmarshal([]byte(values), &post)
	if err != nil {
		return nil, false, err
	}
	return &post, true, nil
}

func SetPostToCache(post *models.Post) error {
	key := GetPostCacheKey(post.ID)
	data, err := json.Marshal(post)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, data, 0).Err()
}

// 定时任务同步缓存到数据库
func SyncCacheToDB() {
	for {

		keys, err := redisClient.Keys(ctx, "post:*").Result()
		if err != nil {
			log.Printf("Error getting keys: %v\n", err)
			time.Sleep(5 * time.Minute)
			continue
		}

		for _, key := range keys {

			val, err := redisClient.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Error getting post data from cache: %v\n", err)
				continue
			}

			var post models.Post
			err = json.Unmarshal([]byte(val), &post)
			if err != nil {
				log.Printf("Error unmarshaling post data: %v\n", err)
				continue
			}

			err = database.DB.Model(&models.Post{}).Where("id =?", post.ID).Update("likes", post.Likes).Error
			if err != nil {
				log.Printf("Error updating database: %v\n", err)
			}
		}

		time.Sleep(5 * time.Minute)
	}
}
