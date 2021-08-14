package databases

import (
	"github.com/go-redis/redis/v8"
)

var (
	TokenQueueClient         *redis.Client
	ResetPasswordQueueClient *redis.Client
)

func (db *DBConfig) TokenRedisConnString() string {
	dns := db.config.TokenRdHost + ":" + db.config.TokenRdPort
	return dns
}
func (db *DBConfig) ResetPasswordRedisConnString() string {
	dns := db.config.ResetpasswordRdHost + ":" + db.config.ResetpasswordRdPort
	return dns
}
