package server

import "github.com/gin-gonic/gin"

// Initialize initiate the server
func Initialize() *gin.Engine {
	r := gin.Default()

	r.GET("/api/local/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.Run()

	// cloudRouter(r, e)
	return r
}
