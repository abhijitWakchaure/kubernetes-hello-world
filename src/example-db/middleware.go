package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	APIKey := c.Request.Header.Get("Authorization")
	if !(APIKey == "DUMMY_API_KEY" || APIKey == "Bearer DUMMY_API_KEY") {
		errorResponse := ResponseError{
			Code:  http.StatusUnauthorized,
			Error: "Invalid API key in auth header, use key 'DUMMY_API_KEY' for testing",
		}
		c.AbortWithStatusJSON(errorResponse.Code, errorResponse)
	}
}
