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

//go:embed version
var version string

func init() {
	log.Printf("Starting DB server [v%s]\n", version)
	var err error
	host, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	meta = Meta{
		Host:        host,
		ServiceType: "db",
	}
	err = os.MkdirAll("./data", 0777)
	if err != nil {
		log.Println("Failed to create data directory due to:", err)
		panic(err)
	}
}

func main() {
	quit := make(chan error)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(authMiddleware)
	router.GET("/", handlerListTasks)
	router.POST("/", handlerInsertTask)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	go func() { quit <- router.Run(port) }()
	log.Printf("DB server [v%s] started on %s:%s \n", version, host, port)
	err := <-quit
	if err != nil {
		panic(err)
	}
}
