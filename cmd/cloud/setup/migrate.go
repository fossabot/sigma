package setup

import (
	"sigma/domain/core/engine"
	"sigma/domain/managing"
)

// Migrate initiate the db connection by getting help from gorm
func Migrate(engine *engine.Engine) {

	if engine.Env.Setting.AutoMigrate {
		managing.Migrate(engine)
	}

}
