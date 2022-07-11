package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()
		logger.Info(
			"incoming request ",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)

	})
	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("hello", func(context *gin.Context) {
		context.String(200, "hello! my master")
	})
	r.Run()
}
