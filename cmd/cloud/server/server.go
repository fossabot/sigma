package server

import (
	"fmt"
	"log"
	"sigma/domain/core/engine"

	"github.com/gin-gonic/gin"
)

// Initialize initiate the server
func Initialize(engine *engine.Engine) *gin.Engine {

	r := gin.Default()

	r.GET("/api/local/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(fmt.Sprintf("%v:%v", engine.Env.Server.ADDR, engine.Env.Server.Port)); err != nil {
		log.Fatalln(err)
	}

	// cloudRouter(r, e)
	return r
}
