package utils

import (
	"CMS/app/models"
	"CMS/config/database"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func GetPostCacheKey(PostID uint) string {
	return fmt.Sprintf("post:%d", PostID)
}

// 从缓存获取点赞值
func GetPostFromCache(PostID uint,c *gin.Context) (int, bool, error) {
	key := GetPostCacheKey(PostID)
	value, err := redisClient.Get(c, key).Int()
	if err != nil {
		if err == redis.Nil {
			return 0, false, nil
		}
		return 0, false, err
	}
	
	return value, true, nil
}
//likes原子自增
func LikesIncr(PostID int,c *gin.Context)error{
	key := GetPostCacheKey(uint(PostID))
	_ ,err:=redisClient.Incr(c,key).Result()
	return err
}
//设置缓存
func SetPostToCache(post *models.Post,c *gin.Context) error {
	key := GetPostCacheKey(post.ID)
	
	return redisClient.Set(c, key, post.Likes,30*time.Minute).Err()
}

// 定时任务同步缓存到数据库
func SyncCacheToDB() {
		ctx:=context.Background()

		keys, err := redisClient.Keys(ctx, "post:*").Result()
		if err != nil {
			log.Printf("Error getting keys: %v\n", err)
			return
		}

		for _, key := range keys {
			post_id:=key[len("post:"):]

			val, err := redisClient.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Error getting post data from cache: %v\n", err)
				return
			}

			err = database.DB.Model(&models.Post{}).Where("id =?", post_id).Update("likes", val).Error
			if err != nil {
				log.Printf("Error updating database: %v\n", err)
				return
			}
		}
}
