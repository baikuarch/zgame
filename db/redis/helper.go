package redis

import "github.com/garyburd/redigo/redis"

// 常用Redis操作命令的封装
// http://redis.io/commands

// KEYS get patten key array
func KEYS(patten string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", patten))
}

// SCAN 获取大量key
func SCAN(patten string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()
	var out []string
	var cursor uint64 = 0xffffff
	isfirst := true
	for cursor != 0 {
		if isfirst {
			cursor = 0
			isfirst = false
		}
		arr, err := conn.Do("SCAN", cursor, "MATCH", patten, "COUNT", 100)
		if err != nil {
			return out, err
		}
		switch arr := arr.(type) {
		case []interface{}:
			cursor, _ = redis.Uint64(arr[0], nil)
			it, _ := redis.Strings(arr[1], nil)
			out = append(out, it...)
		}
	}
	out = ArrayDuplice(out)
	return out, nil
}

// DEL delete k-v
func DEL(key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("DEL", key))
}

// DELALL delete key array
func DELALL(key []string) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	arr := make([]interface{}, len(key))
	for i, v := range key {
		arr[i] = v
	}
	return redis.Int(conn.Do("DEL", arr...))
}

// GET get k-v
func GET(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// SET set k-v
func SET(key string, value string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("SET", key, value))
}

// SETEX set k-v expire seconds
func SETEX(key string, sec int, value string) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SETEX", key, sec, value))
}

// EXPIRE set key expire seconds
func EXPIRE(key string, sec int64) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("EXPIRE", key, sec))
}

// HGETALL get map of key
func HGETALL(key string) (map[string]string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}

// HGET get value of key-field
func HGET(key string, field string) (string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET", key, field))
}

// HSET set value of key-field
func HSET(key string, field string, value string) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("HSET", key, field, value))
}
