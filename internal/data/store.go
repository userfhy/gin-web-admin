package data

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// Store 封装基础数据访问依赖，便于在 service 层注入。
type Store struct {
	db        *gorm.DB
	redisPool *redis.Pool
}

func NewStore(db *gorm.DB, redisPool *redis.Pool) *Store {
	return &Store{db: db, redisPool: redisPool}
}

func (s *Store) DB() *gorm.DB {
	return s.db
}

func (s *Store) Redis() *redis.Pool {
	return s.redisPool
}
