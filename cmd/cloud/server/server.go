package server

import (
	"fmt"
	"log"
	"sigma/domain/core"

	envEngine "github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
)

// Initialize initiate the server
func Initialize() *gin.Engine {
	var engine core.Engine
	if err := envEngine.Parse(&engine.Env); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()

	r.GET("/api/local/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(fmt.Sprintf("%v:%v", "localhost", 8183)); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Cloud Server started \n\n")

	// cloudRouter(r, e)
	return r
}
