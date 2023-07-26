package db

import "github.com/baikuarch/zgame/db/redis"

type Db struct {
}

const (
	ENUM_Redis = iota
	ENUM_MySQL
)

// InitDb 初始化各类数据库连接
func (db *Db) InitDb(dbArray []int8) {
	if dbArray != nil && len(dbArray) > 0 {
		for i := 0; i < len(dbArray); i++ {
			dbType := dbArray[i]
			switch dbType {
			case ENUM_Redis:
				_ = redis.NewRedis()
				break

			}
		}
	}
}

// CloseDb 关闭各类数据库连接
func (db *Db) CloseDb() {
	//redisConn.Close()
}

func NewDb() *Db {
	return &Db{}
}
