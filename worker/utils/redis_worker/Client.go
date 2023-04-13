package redis_worker

import (
	"fmt"
	"time"
	"context"
	"sync"
	"strconv"
	redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ConnectToRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	// There is no error because go-redis automatically reconnects on error.
	pubsubCounter := client.Subscribe(ctx, "__keyspace@0__:counter")
	// Close the subscription when we are done.
	defer pubsubCounter.Close()

	// There is no error because go-redis automatically reconnects on error.
	pubsubNoti := client.Subscribe(ctx, "__keyspace@0__:notifyCounter")
	// Close the subscription when we are done.
	defer pubsubNoti.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Monitoring on counter")
		for {
			msg, err := pubsubCounter.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}
	
			fmt.Println(msg.Channel, msg.Payload)
	
			if msg.Payload == "expired" {
				client.Set(ctx, "counter", 10, 30 * time.Second)
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Monitoring on notifyCounter")
		for {
			msg, err := pubsubNoti.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}
	
			fmt.Println(msg.Channel, msg.Payload)
	
			if msg.Payload == "set" {
				ttl, _ := client.TTL(ctx, "counter").Result()
				counterValueStr, _  := client.Get(ctx, "counter").Result()
				counterValue, _ := strconv.Atoi(counterValueStr)
				client.Set(ctx, "counter", counterValue - 1, ttl)
			}
		}
	}()
	wg.Wait()
}