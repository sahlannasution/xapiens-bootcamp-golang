package controllers

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func CheckRedis(key string) (bool, []byte) {
	// make a pool to connect to redis
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10)
	//
	pool.MaxActive = 10

	// Get connection
	conn := pool.Get()
	defer conn.Close()

	// Finding Data with key
	findData, _ := redis.Bytes(conn.Do("GET", key))
	if findData != nil { // ketika ada datanya di redis
		log.Println("Data Found!")
		log.Println(string(findData))
		return true, findData
	}
	return false, findData
}

func SetRedis(key string, value string) {
	// make a pool to connect to redis
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10)
	//
	pool.MaxActive = 10

	// Get connection
	conn := pool.Get()
	defer conn.Close()

	// Finding Data with key
	conn.Do("SETEX", key, 30, string(value))
}

// func GetRedis(key string) string {
// 	// make a pool to connect to redis
// 	pool := redis.NewPool(func() (redis.Conn, error) {
// 		return redis.Dial("tcp", "localhost:6379")
// 	}, 10)
// 	//
// 	pool.MaxActive = 10

// 	// Get connection
// 	conn := pool.Get()
// 	defer conn.Close()

// 	// Finding Data with key
// 	getData, _ := redis.Bytes(conn.Do("GET", key))
// 	if getData != nil { // ketika ada datanya di redis
// 		log.Println("Data Found!")
// 		log.Println(string(getData))

// 	}
// 	return string(getData)
// }
