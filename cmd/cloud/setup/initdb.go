package setup

import (
	"log"
	"sigma/domain/core"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// InitDB initiate the db connection by getting help from gorm
func InitDB(engine *core.Engine) {
	var err error
	engine.DB, err = gorm.Open(engine.Env.Database.Data.Type, engine.Env.Database.Data.DSN)
	if err != nil {
		log.Fatalln(err.Error())
	}

	engine.DB.LogMode(true)

	if gin.IsDebugging() {
		engine.DB.LogMode(true)
	}

}
