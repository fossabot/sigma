package core

import "github.com/jinzhu/gorm"

// Engine to keep all database connections and
// logs configuration and environments and etc
type Engine struct {
	DB *gorm.DB
	// ActivityDB   *gorm.DB
	// ServerLog    *logrus.Logger
	// APILog       *logrus.Logger
	Env Environment
}
