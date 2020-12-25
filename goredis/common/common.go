package common

import "github.com/go-redis/redis"

const (
	OneWeekInSeconds       = 7 * 86400
	VoteScore              = 432
	ArticlesPerPage  int64 = 25
)

var rdb *redis.Client

// 初始化连接
func InitClient() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.14.137:16379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
