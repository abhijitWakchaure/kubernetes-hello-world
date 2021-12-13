package main

import (
	_ "embed"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var meta Meta
var host string
var apiKey, dbServiceURL string

//go:embed version
var version string

func init() {
	log.Printf("Starting app server [v%s]\n", version)
	var err error
	host, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	meta = Meta{
		Host:        host,
		ServiceType: "app",
	}
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		panic("API_KEY is not set")
	}
	dbServiceURL = os.Getenv("DB_SERVICE_URL")
	if dbServiceURL == "" {
		panic("DB_SERVICE_URL is not set")
	}
	log.Println("DB_SERVICE_URL set to:", dbServiceURL)
}

func main() {
	quit := make(chan error)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(RespHeaderMiddleware)
	router.GET("/", ReverseProxy)
	router.POST("/", ReverseProxy)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	go func() { quit <- router.Run(port) }()
	log.Printf("App server [v%s] started on %s:%s \n", version, host, port)
	err := <-quit
	if err != nil {
		panic(err)
	}
}
