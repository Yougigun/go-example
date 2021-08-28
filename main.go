package main

import "github.com/gin-gonic/gin"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router := gin.New()
	// router.Use(gin.Recovery())
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.GET("/panic", func(c *gin.Context) {
		panic("testing panic")
	})
	router.GET("/ok", func(c *gin.Context) {
		c.JSON(200,"ok")
	})
	router.Run(":8080")
}