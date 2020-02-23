package managing

import (
	"sigma/domain/core/engine"
	"sigma/domain/managing/model"
)

// Migrate is called for creating tables, indexes and etc
func Migrate(e *engine.Engine) {
	e.DB.AutoMigrate(&model.Role{})
}
