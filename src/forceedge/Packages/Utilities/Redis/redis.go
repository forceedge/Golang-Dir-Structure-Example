package redis

import "fmt"
import "github.com/garyburd/redigo/redis"
import "time"

func Connect() *redis.Pool {
	server := "6379"

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			return c, err
		},
	}
}

func Store(doc string, value string) interface{} {
	conn := Connect().Get()
	defer conn.Close()

	// Increment the message ID.
	msgId, msgIdErr := conn.Do("INCR", "messageId")
	if msgIdErr != nil {
		panic(msgIdErr)
	}

	_, err := conn.Do("SET", doc, msgId, value)
	if err != nil {
		panic(err)
	}

	return msgId
}

func Get(doc string, identifier int) string {
	conn := Connect().Get()
	defer conn.Close()

	exists, existsErr := redis.Int(conn.Do("HEXISTS", doc, identifier))
	if existsErr != nil {
		panic(existsErr)
	}

	if exists != 1 {
		return "message with this ID does not exist"
	}

	storedValue, err := redis.String(conn.Do("GET", identifier))

	if err != nil {
		fmt.Println("key not found")
	}

	return storedValue
}
