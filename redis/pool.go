package redis

import (
	"github.com/mediocregopher/radix.v2/pool"
)

var DredisHost = "localhost"
var DredisPorts = []int{6380}
var DredisPoolSize = 10

var DredisPool *pool.Pool

func init() {
	DredisPool = createDredisPool()
}

func createDredisPool() *pool.Pool {
	p, err := pool.New("tcp", DredisHost+":"+"6380", 10)
	if err != nil {
		// handle error
	}
	return p
}
