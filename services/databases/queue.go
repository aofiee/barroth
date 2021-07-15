package databases

import (
	"github.com/go-redis/redis/v8"
)

var (
	QueueClient *redis.Client
)

func (db *DBConfig) RedisConnString() string {
	dns := db.config.RdHost + ":" + db.config.RdPort
	return dns
}
