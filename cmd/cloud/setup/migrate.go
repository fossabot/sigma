package setup

import (
	"sigma/domain/core"
	"sigma/domain/managing"
)

// Migrate initiate the db connection by getting help from gorm
func Migrate(engine *core.Engine) {

	if engine.Env.Setting.AutoMigrate {
		managing.Migrate(engine)
	}

}
