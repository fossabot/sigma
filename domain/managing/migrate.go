package managing

import (
	"sigma/domain/core"
	"sigma/domain/managing/model"
)

// Migrate is called for creating tables, indexes and etc
func Migrate(e *core.Engine) {
	e.DB.AutoMigrate(&model.Role{})
}
