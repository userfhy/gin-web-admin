package gredis

import (
	"encoding/json"
	"fmt"
	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/setting"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}

			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", setting.RedisSetting.DB); err != nil {
				_ = c.Close()
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	logging.Printf("Redis connected %s DB: %d", setting.RedisSetting.Host, setting.RedisSetting.DB)

	TestConnection()
}

func TestConnection() {
	conn := RedisConn.Get()
	defer closeConn(conn)

	res, err := conn.Do("PING")
	if err != nil {
		panic(err)
	}
	logging.Println(res)
}

// Set a key/value
func Set(key string, data any, time int) error {
	conn := RedisConn.Get()
	defer closeConn(conn)

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer closeConn(conn)

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer closeConn(conn)

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer closeConn(conn)

	return redis.Bool(conn.Do("DEL", key))
}

// DeleteKeys delete multiple keys, returns number of removed keys
func DeleteKeys(keys ...string) (int, error) {
	if RedisConn == nil {
		return 0, fmt.Errorf("redis not initialized")
	}
	if len(keys) == 0 {
		return 0, nil
	}
	args := make([]any, 0, len(keys))
	for _, key := range keys {
		args = append(args, key)
	}
	conn := RedisConn.Get()
	defer closeConn(conn)
	return redis.Int(conn.Do("DEL", args...))
}

// DeleteByPrefix 删除指定前缀的所有 key
func DeleteByPrefix(prefix string) (int, error) {
	if RedisConn == nil {
		return 0, fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)

	cursor := 0
	total := 0
	match := prefix + "*"

	for {
		values, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", match, "COUNT", 100))
		if err != nil {
			return total, err
		}
		var keys []string
		if _, err := redis.Scan(values, &cursor, &keys); err != nil {
			return total, err
		}
		if len(keys) > 0 {
			args := make([]any, 0, len(keys))
			for _, k := range keys {
				args = append(args, k)
			}
			deleted, err := redis.Int(conn.Do("DEL", args...))
			if err != nil {
				return total, err
			}
			total += deleted
		}
		if cursor == 0 {
			break
		}
	}
	return total, nil
}

// KeysByPrefix 获取指定前缀的所有 key
func KeysByPrefix(prefix string) ([]string, error) {
	if RedisConn == nil {
		return nil, fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)

	cursor := 0
	match := prefix + "*"
	keys := make([]string, 0)

	for {
		values, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", match, "COUNT", 100))
		if err != nil {
			return keys, err
		}
		var batch []string
		if _, err := redis.Scan(values, &cursor, &batch); err != nil {
			return keys, err
		}
		keys = append(keys, batch...)
		if cursor == 0 {
			break
		}
	}

	return keys, nil
}

// DeleteByPrefixAsync 异步删除指定前缀，避免阻塞调用方
func DeleteByPrefixAsync(prefix string) {
	if RedisConn == nil {
		return
	}
	go func() {
		if _, err := DeleteByPrefix(prefix); err != nil {
			logging.Warnf("redis delete prefix %s failed: %v", prefix, err)
		}
	}()
}

// DeleteKeysAsync 异步删除
func DeleteKeysAsync(keys ...string) {
	if len(keys) == 0 || RedisConn == nil {
		return
	}
	go func() {
		if _, err := DeleteKeys(keys...); err != nil {
			logging.Warnf("redis delete keys failed: %v", err)
		}
	}()
}

// SetWithTTL 写入字符串并设置过期时间
func SetWithTTL(key string, value any, ttl time.Duration) error {
	if RedisConn == nil {
		return fmt.Errorf("redis not initialized")
	}
	if ttl <= 0 {
		ttl = time.Minute
	}
	conn := RedisConn.Get()
	defer closeConn(conn)

	seconds := int(ttl.Seconds())
	if seconds <= 0 {
		seconds = 60
	}

	_, err := conn.Do("SETEX", key, seconds, value)
	return err
}

// SetWithTTLAsync 异步写入
func SetWithTTLAsync(key string, value any, ttl time.Duration) {
	if RedisConn == nil {
		return
	}
	go func() {
		if err := SetWithTTL(key, value, ttl); err != nil {
			logging.Warnf("redis set with ttl failed: %v", err)
		}
	}()
}

// ExistsKey 判断键是否存在
func ExistsKey(key string) (bool, error) {
	if RedisConn == nil {
		return false, fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return exists, nil
}

// SetJSON 写入 JSON 数据，可选 TTL
func SetJSON(key string, value any, ttl time.Duration) error {
	if RedisConn == nil {
		return fmt.Errorf("redis not initialized")
	}
	payload, err := json.Marshal(value)
	if err != nil {
		return err
	}
	conn := RedisConn.Get()
	defer closeConn(conn)

	if ttl > 0 {
		_, err = conn.Do("SETEX", key, int(ttl.Seconds()), payload)
	} else {
		_, err = conn.Do("SET", key, payload)
	}
	return err
}

// GetJSON 读取 JSON 数据，返回是否命中
func GetJSON(key string, dest any) (bool, error) {
	if RedisConn == nil {
		return false, fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)

	data, err := redis.Bytes(conn.Do("GET", key))
	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(data, dest); err != nil {
		return false, err
	}
	return true, nil
}

// SetJSONAsync 异步写入 JSON
func SetJSONAsync(key string, value any, ttl time.Duration) {
	if RedisConn == nil {
		return
	}
	go func() {
		if err := SetJSON(key, value, ttl); err != nil {
			logging.Warnf("redis set json failed: %v", err)
		}
	}()
}

// SetString 写入字符串，可选 TTL（<=0 则不过期）
func SetString(key, value string, ttl time.Duration) error {
	if RedisConn == nil {
		return fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)

	if ttl > 0 {
		_, err := conn.Do("SETEX", key, int(ttl.Seconds()), value)
		return err
	}
	_, err := conn.Do("SET", key, value)
	return err
}

// GetString 读取字符串
func GetString(key string) (string, error) {
	if RedisConn == nil {
		return "", fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)
	return redis.String(conn.Do("GET", key))
}

// Expire 设置过期时间
func Expire(key string, ttl time.Duration) error {
	if RedisConn == nil {
		return fmt.Errorf("redis not initialized")
	}
	if ttl <= 0 {
		return fmt.Errorf("ttl must be positive")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)
	_, err := conn.Do("EXPIRE", key, int(ttl.Seconds()))
	return err
}

// HSet 写入哈希字段
func HSet(key, field string, value any) error {
	if RedisConn == nil {
		return fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)
	_, err := conn.Do("HSET", key, field, value)
	return err
}

// HGetString 读取哈希字段（字符串）
func HGetString(key, field string) (string, error) {
	if RedisConn == nil {
		return "", fmt.Errorf("redis not initialized")
	}
	conn := RedisConn.Get()
	defer closeConn(conn)
	return redis.String(conn.Do("HGET", key, field))
}

// HDel 删除哈希字段
func HDel(key string, fields ...string) (int, error) {
	if RedisConn == nil {
		return 0, fmt.Errorf("redis not initialized")
	}
	if len(fields) == 0 {
		return 0, nil
	}
	args := make([]any, 0, len(fields)+1)
	args = append(args, key)
	for _, f := range fields {
		args = append(args, f)
	}
	conn := RedisConn.Get()
	defer closeConn(conn)
	return redis.Int(conn.Do("HDEL", args...))
}

func closeConn(conn redis.Conn) {
	if err := conn.Close(); err != nil {
		logging.Warnf("redis conn close failed: %v", err)
	}
}
