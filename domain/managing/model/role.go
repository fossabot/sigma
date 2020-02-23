package model

import "sigma/internal/types"

// Role model
type Role struct {
	types.FixedCol
	Name        string `gorm:"not null;unique" json:"name,omitempty" binding:"required"`
	Resources   string `gorm:"type:text" json:"resources,omitempty" binding:"-"`
	Description string `json:"description,omitempty" binding:"-"`
}
