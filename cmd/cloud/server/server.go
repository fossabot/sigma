package server

import (
	"fmt"
	"log"
	"sigma/domain/core/engine"
	"sigma/domain/managing"

	envEngine "github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Initialize initiate the server
func Initialize() *gin.Engine {
	var err error
	var engine engine.Engine
	if err = envEngine.Parse(&engine.Env); err != nil {
		log.Fatalln(err)
	}

	engine.DB, err = gorm.Open(engine.Env.Database.Data.Type, engine.Env.Database.Data.DSN)
	if err != nil {
		log.Fatalln(err.Error())
	}

	engine.DB.LogMode(true)

	if gin.IsDebugging() {
		engine.DB.LogMode(true)
	}

	managing.Migrate()

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
