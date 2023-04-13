package main

import (
	. "ratelimiterworker.com/m/v2/redis_router"
	"github.com/gin-gonic/gin"
)

// main Entry point
func main() {
    router := gin.Default()
    router.GET("/counter", GetCounter)

    router.Run("localhost:8080")
}