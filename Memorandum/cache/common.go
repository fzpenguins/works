package cache

import (
	"Memorandum/config"
	"Memorandum/model"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func Redis() {
	db, _ := strconv.ParseUint(config.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}

func TaskView(Task model.Task) uint64 {
	// 增加点击数
	countStr, _ := RedisClient.Get(TaskViewKey(Task.Id)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func AddView(Task *model.Task) {
	RedisClient.Incr(TaskViewKey(Task.Id))                      // 增加视频点击数
	RedisClient.ZIncrBy(RankKey, 1, strconv.Itoa(int(Task.Id))) // 增加排行点击数
}
