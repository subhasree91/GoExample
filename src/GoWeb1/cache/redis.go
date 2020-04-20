package cache

import (
	"encoding/json"
	"fmt"
	"log"

	// "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
)

func redisConnec() redis.Conn {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.

	return conn

	// Send our command across the connection. The first parameter to
	// Do() is always the name of the Redis command (in this example
	// HMSET), optionally followed by any necessary arguments (in this
	// example the key, followed by the various hash fields and values).
	// _, err = conn.Do("HMSET", "album:2", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	// if err != nil {
	// 	fmt.Println("Failed HMSET")
	// 	log.Fatal(err)
	// }

	// title, err := redis.String(conn.Do("HGET", "album:2", "title"))
	// if err != nil {
	// 	fmt.Println("failed for HGET")
	// 	log.Fatal(err)
	// }
	// fmt.Printf(title)

}

func RedisSet(x10 string, abc interface{}) {
	fmt.Println("redis set")
	conn := redisConnec()

	defer conn.Close()

	_, err := conn.Do("HMSET", x10, "event", abc)
	if err != nil {
		fmt.Println("Failed HMSET")
		log.Fatal(err)
	}

	fmt.Println("abc=", abc)

}

func RedisGet(x10 string) interface{} {

	conn := redisConnec()

	defer conn.Close()
	title, err := redis.String(conn.Do("HGET", x10, "event"))
	if err != nil {
		fmt.Println("failed for HGET")
		return nil
	}
	var output interface{}
	err = json.Unmarshal([]byte(title), &output)
	if err != nil {
		fmt.Println("err in unmarshal", err)
		fmt.Println("title", title)
	}
	fmt.Println("redis get", output)
	return output
}

// func redisCache() {
// 	var pool = newPool()
// 	c := pool.Get()
// 	defer c.Close()
// }
// func newPool() *redis.PoolStats {
// 	return &redis.Pool{
// 		MaxIdle:   80,
// 		MaxActive: 12000, // max number of connections
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.Dial("tcp", ":6379")
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			return c, err
// 		},
// 	}
// }
