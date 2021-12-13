package main

import (
	_ "embed"
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
	initLogger()
	logger.Infof("Starting app server [v%s]", version)
	var err error
	host, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	logger.Debugf("Hostname: %s", host)
	meta = Meta{
		Host:        host,
		ServiceType: "app",
	}
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		panic("API_KEY is not set")
	}
	logger.Debugf("API_KEY is set to: %s", apiKey)
	dbServiceURL = os.Getenv("DB_SERVICE_URL")
	if dbServiceURL == "" {
		panic("DB_SERVICE_URL is not set")
	}
	logger.Infof("DB_SERVICE_URL is set to: %s", dbServiceURL)
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
	logger.Infof("App server [v%s] started on %s:%s \n", version, host, port)
	err := <-quit
	if err != nil {
		panic(err)
	}
}
