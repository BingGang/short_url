package redis

import (
	"github.com/garyburd/redigo/redis"
)

func  (this *Redis)EXPIRE(k string, v interface{}) (bool, error) {
	result, err := redis.Int(this.Do("EXPIRE", k, v))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result != 0, nil
}
