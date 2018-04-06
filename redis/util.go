package redis

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
)

func Ping() error {

	conn, err := DredisPool.Get()
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	defer DredisPool.Put(conn)

	result := conn.Cmd("PING")
	if result.Err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", result.Err)
	}
	return nil
}

func GetConnection() *redis.Client {
	conn, err := DredisPool.Get()
	if err != nil {

	}
	defer DredisPool.Put(conn)
	return conn
}
