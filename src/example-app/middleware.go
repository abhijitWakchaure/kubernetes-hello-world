package main

import "github.com/gin-gonic/gin"

// RespHeaderMiddleware ...
func RespHeaderMiddleware(c *gin.Context) {
	c.Writer.Header().Set("X-App-ServerId", meta.Host)
	c.Next()
}
