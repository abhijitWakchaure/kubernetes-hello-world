package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// ReverseProxy ...
func ReverseProxy(c *gin.Context) {
	director := func(req *http.Request) {
		backendURL, _ := url.Parse(dbServiceURL)
		req.URL = backendURL
		req.Header.Set("Authorization", apiKey)
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
