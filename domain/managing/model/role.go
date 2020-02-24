package model

import (
	"fmt"
	"sigma/internal/types"
)

// Role model
type Role struct {
	types.FixedCol
	Name        string `gorm:"not null;unique" json:"name,omitempty" binding:"required"`
	Resources   string `gorm:"type:text" json:"resources,omitempty" binding:"-"`
	Description string `json:"description,omitempty" binding:"-"`
}

// SearchPattern is used for generate query for finding ter  among specific columns
func (p *Role) SearchPattern(str string) string {
	var pattern = `(roles.name LIKE '%%%[1]v%%' OR
		roles.resources LIKE '%%%[1]v%%')`
	return fmt.Sprintf(pattern, str)
}
