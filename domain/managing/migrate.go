package managing

import (
	"fmt"
	"sigma/domain/core/engine"
	"sigma/domain/managing/model"
)

// import

// Migrate is called for creating tables, indexes and etc
func Migrate(e *engine.Engine) {
	e.DB.AutoMigrate(&model.Role{})

}
