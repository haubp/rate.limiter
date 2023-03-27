package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// There is no error because go-redis automatically reconnects on error.
	pubsub := client.Subscribe("__keyspace@0__:counter")

	// Close the subscription when we are done.
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
	
		fmt.Println(msg.Channel, msg.Payload)
	}
}